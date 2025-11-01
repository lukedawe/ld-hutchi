// Response DTOs.
// Mirrors of what can be found in api/handlers/dtos/responses/

import * as z from "zod";

export const BreedResponseZod = z.object({
    id: z.number(),
    name: z.string(),
    category_id: z.number(),
});

export const CategoryResponseZod = z.object({
    id: z.number(),
    name: z.string(),
    breeds: z.array(BreedResponseZod)
});

export const ErrorResponseZod = z.object({
    code: z.string(),
    message: z.string(),
})

export type CategoryResponse = z.infer<typeof CategoryResponseZod>;
export type BreedResponse = z.infer<typeof BreedResponseZod>;
export type ErrorResponse = z.infer<typeof ErrorResponseZod>;