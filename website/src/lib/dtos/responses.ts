// Response DTOs.

export interface BreedResponse {
    id: number;
    name: string;
}

export interface CategoryResponse {
    id: number;
    name: string;
    breeds: BreedResponse[];
};

function isObject(v: unknown): v is Record<string, unknown> {
    return typeof v === 'object' && v !== null;
}

function assertIsNumber(v: unknown, name = 'value'): number {
    if (typeof v !== 'number' || Number.isNaN(v)) throw new Error(`${name} is not a number`);
    return v;
}

function assertIsString(v: unknown, name = 'value'): string {
    if (typeof v !== 'string') throw new Error(`${name} is not a string`);
    return v;
}

export function jsonToCategoryResponse(json: any): CategoryResponse {
    if (!isObject(json)) throw new Error('expected object for CategoryResponse');

    const id = assertIsNumber((json as any).id, 'id');
    const name = assertIsString((json as any).name, 'name');

    const breedsRaw = (json as any).breeds;
    if (!Array.isArray(breedsRaw)) throw new Error('breeds is not an array');

    const breeds: BreedResponse[] = breedsRaw.map((b, idx) => {
        try {
            return jsonToBreedResponse(b);
        } catch (err) {
            throw new Error(`invalid breed at index ${idx}: ${(err as Error).message}`);
        }
    });

    return { id, name, breeds };
}

export function jsonToBreedResponse(json: any): BreedResponse {
    if (!isObject(json)) throw new Error('expected object for BreedResponse');

    const id = assertIsNumber((json as any).id, 'id');
    const name = assertIsString((json as any).name, 'name');

    return { id, name };
}