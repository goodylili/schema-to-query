
import * as fs from 'fs';
import * as path from 'path';
import {
  buildClientSchema,
  isScalarType,
  isObjectType,
  isListType,
  isNonNullType,
  isInterfaceType,
  isUnionType,
  getIntrospectionQuery,
} from 'graphql';
import type {
  IntrospectionQuery,
  GraphQLSchema,
  GraphQLField,
  GraphQLType,
} from 'graphql';

// A safe maximum depth to prevent OOM/StackOverflow on large schemas.
// 3 is generally deep enough for almost all use cases while being safer than "infinite".
export const DEFAULT_MAX_DEPTH = 3;

function getUnwrappedType(type: GraphQLType): GraphQLType {
  if (isNonNullType(type)) {
    return getUnwrappedType(type.ofType);
  }
  if (isListType(type)) {
    return getUnwrappedType(type.ofType);
  }
  return type;
}

function hasValidSelection(type: GraphQLType, maxDepth: number, visitedTypes: Set<string>, depth: number): boolean {
    if (depth > maxDepth) {
        return false;
    }

    const unwrappedType = getUnwrappedType(type);

    if (isScalarType(unwrappedType)) {
        // Scalars are "leaves" in the selection tree, but this function checks if we can *select* from them.
        // Wait, if we are at a field F of type Scalar, we definitely include it.
        // But the recursion calls hasValidSelection(fieldType).
        // If fieldType is scalar, it means we found a leaf.
        return true;
    }

    const typeName = 'name' in unwrappedType ? (unwrappedType as any).name : '';
    if (typeName && visitedTypes.has(typeName)) {
        return false;
    }

    const newVisitedTypes = new Set(visitedTypes);
    if (typeName) {
        newVisitedTypes.add(typeName);
    }

    if (isObjectType(unwrappedType) || isInterfaceType(unwrappedType)) {
        const fields = unwrappedType.getFields();
        for (const field of Object.values(fields)) {
            const fieldType = getUnwrappedType(field.type);
            if (isScalarType(fieldType)) {
                return true;
            }
            if (hasValidSelection(field.type, maxDepth, newVisitedTypes, depth + 1)) {
                return true;
            }
        }
        return false;
    }

    if (isUnionType(unwrappedType)) {
        const types = unwrappedType.getTypes();
        for (const t of types) {
            if (hasValidSelection(t, maxDepth, newVisitedTypes, depth + 1)) {
                return true;
            }
        }
        return false;
    }

    return false;
}

async function writeFieldSelection(
    stream: fs.WriteStream,
    type: GraphQLType,
    maxDepth: number,
    visitedTypes: Set<string>,
    depth: number,
    indent: string
): Promise<void> {
    if (depth > maxDepth) {
        return;
    }

    const unwrappedType = getUnwrappedType(type);
    const typeName = 'name' in unwrappedType ? (unwrappedType as any).name : '';

    if (typeName && visitedTypes.has(typeName)) {
        return;
    }

    const newVisitedTypes = new Set(visitedTypes);
    if (typeName) {
        newVisitedTypes.add(typeName);
    }

    const write = async (chunk: string) => {
        if (!stream.write(chunk)) {
            await new Promise<void>((resolve) => stream.once('drain', resolve));
        }
    };

    if (isObjectType(unwrappedType) || isInterfaceType(unwrappedType)) {
        const fields = unwrappedType.getFields();
        for (const field of Object.values(fields)) {
            const fieldType = getUnwrappedType(field.type);
            
            if (isScalarType(fieldType)) {
                await write(`${indent}${field.name}\n`);
                continue;
            }

            if (hasValidSelection(field.type, maxDepth, newVisitedTypes, depth + 1)) {
                await write(`${indent}${field.name} {\n`);
                await writeFieldSelection(stream, field.type, maxDepth, newVisitedTypes, depth + 1, indent + '  ');
                await write(`${indent}}\n`);
            }
        }
    } else if (isUnionType(unwrappedType)) {
        const types = unwrappedType.getTypes();
        for (const t of types) {
            if (hasValidSelection(t, maxDepth, newVisitedTypes, depth + 1)) {
                await write(`${indent}... on ${t.name} {\n`);
                await writeFieldSelection(stream, t, maxDepth, newVisitedTypes, depth + 1, indent + '  ');
                await write(`${indent}}\n`);
            }
        }
    }
}

async function writeOperation(
    stream: fs.WriteStream,
    operationType: 'query' | 'mutation',
    fieldName: string,
    field: GraphQLField<any, any>,
    maxDepth: number
): Promise<void> {
    const args = field.args;
    let queryName = `${fieldName.charAt(0).toUpperCase() + fieldName.slice(1)}`;
    if (operationType === 'mutation') queryName = 'Mutate' + queryName;
    else queryName = 'Get' + queryName;

    let varDefs = '';
    let argsUsage = '';

    if (args.length > 0) {
        const vars = args.map(arg => {
            const typeStr = arg.type.toString();
            return `$${arg.name}: ${typeStr}`;
        });
        varDefs = `(${vars.join(', ')})`;

        const usages = args.map(arg => {
            return `${arg.name}: $${arg.name}`;
        });
        argsUsage = `(${usages.join(', ')})`;
    }

    const write = async (chunk: string) => {
        if (!stream.write(chunk)) {
            await new Promise<void>((resolve) => stream.once('drain', resolve));
        }
    };

    if (hasValidSelection(field.type, maxDepth, new Set(), 0)) {
        await write(`${operationType} ${queryName}${varDefs} {\n`);
        await write(`  ${fieldName}${argsUsage} {\n`);
        await writeFieldSelection(stream, field.type, maxDepth, new Set(), 0, '    ');
        await write(`  }\n`);
        await write(`}\n\n`);
    } else {
        // Even if no selection (unlikely for root fields usually), we might want to print something or skip.
        // If it returns scalar (unlikely for root query/mutation), we should print it.
        const unwrapped = getUnwrappedType(field.type);
        if (isScalarType(unwrapped)) {
             await write(`${operationType} ${queryName}${varDefs} {\n`);
             await write(`  ${fieldName}${argsUsage}\n`);
             await write(`}\n\n`);
        }
    }
}


async function fetchSchema(endpoint: string): Promise<GraphQLSchema> {
    console.log(`Fetching schema from ${endpoint}...`);
    
    const introspectionQuery = getIntrospectionQuery();

    const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ query: introspectionQuery }),
    });

    if (!response.ok) {
        throw new Error(`Failed to fetch schema: ${response.status} ${response.statusText}`);
    }

    const result = await response.json();

    if (result.errors) {
        throw new Error(`Schema introspection errors: ${JSON.stringify(result.errors)}`);
    }

    const introspectionData = result.data;
    return buildClientSchema(introspectionData as IntrospectionQuery);
}

export async function generate(endpoint: string, outputPath: string, maxDepth: number = DEFAULT_MAX_DEPTH) {
    // If maxDepth is 0, we treat it as "Infinity" (relying on cycle detection).
    // However, to prevent stack overflows, we still need a theoretical limit.
    // 100 is deep enough to cover any realistic non-cyclic schema structure.
    const effectiveMaxDepth = maxDepth === 0 ? 100 : maxDepth;

    try {
        const schema = await fetchSchema(endpoint);

        const queryType = schema.getQueryType();
        const mutationType = schema.getMutationType();

        const stream = fs.createWriteStream(outputPath);
        
        // Helper to handle backpressure is now inside writeOperation/writeFieldSelection or we can reuse it if we pass it.
        // But writeOperation takes stream and handles writes internally.

        if (queryType) {
            const fields = queryType.getFields();
            for (const [fieldName, field] of Object.entries(fields)) {
                await writeOperation(stream, 'query', fieldName, field, effectiveMaxDepth);
            }
        }

        if (mutationType) {
            const fields = mutationType.getFields();
            for (const [fieldName, field] of Object.entries(fields)) {
                await writeOperation(stream, 'mutation', fieldName, field, effectiveMaxDepth);
            }
        }

        stream.end();

        await new Promise<void>((resolve, reject) => {
            stream.on('finish', resolve);
            stream.on('error', reject);
        });
        
        console.log(`Generated queries at ${outputPath}`);
    } catch (error: any) {
        // Re-throw specific errors for the caller to handle
        if (error instanceof RangeError || error.message?.includes('Invalid string length')) {
            throw new Error('OUTPUT_TOO_LARGE');
        }
        throw error;
    }
}
