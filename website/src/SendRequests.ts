import type { ZodObject } from "zod";
import { ErrorResponseZod, type ErrorResponse } from "./lib/dtos/responses";

const applicationJson = 'application/json'

export async function GET<T>(uri: string, zod: ZodObject): Promise<T | ErrorResponse> {
    return new Promise(async (resolve, reject) => {
        const res = await fetch(uri, {
            method: 'GET',
        });

        if (!res.ok) {
            const errorResponse = ErrorResponseZod.safeParse(await res.json())
            if (errorResponse.success) {
                resolve(errorResponse.data as ErrorResponse);
                return;
            }
            else reject(new Error("Unexpected response to request."));

            return;
        }

        const responseJson = await res.json();
        const response = zod.safeParse(responseJson);
        if(response.success) {
            return response.data as T;
        }
        else {
            reject(new Error("Unexpected response to request"));
        }
    })
}

export function DELETE(uri: string): Promise<void | ErrorResponse> {
    return new Promise(async (resolve, reject) => {
        const res = await fetch(uri, {
            method: 'DELETE',
        });

        if (!res.ok) {
            const errorResponse = ErrorResponseZod.safeParse(await res.json())
            if (errorResponse.success) {
                resolve(errorResponse.data as ErrorResponse);
                return;
            }
            else reject(new Error("Unexpected response to request."));

            return;
        }

        resolve();
    })
}

export function POST(uri: string, payload: any): Promise<Response> {
    return fetch(uri, {
        method: 'POST',
        headers: { 'Content-Type': applicationJson },
        body: JSON.stringify(payload),
    });
}

export function PATCH(uri: string, payload: any): Promise<Response> {
    return fetch(uri, {
        method: 'PATCH',
        headers: { 'Content-Type': applicationJson },
        body: JSON.stringify(payload),
    });
}

export function PUT(uri: string, payload: any): Promise<Response> {
    return fetch(uri, {
        method: 'PUT',
        headers: { 'Content-Type': applicationJson },
        body: JSON.stringify(payload),
    });
}