// ...new file...
export interface BreedNameRequiredJson {
  name: string;
}

export interface BreedIdRequiredUri {
  id: number;
}

export interface BreedIdOptionalUri {
  id?: number;
}

export interface BreedCategoryIdRequiredJson {
  category_id: number;
}

/** Request payload to add a breed */
export interface AddBreed {
  name: string;
  category_id: number;
}

/** Alias types mirroring Go DTO names */
export type GetBreed = BreedIdRequiredUri;
export type PutBreedUri = BreedIdOptionalUri;
export type PatchBreedUri = BreedIdOptionalUri;
export type PatchBreedBody = BreedNameRequiredJson;
export type DeleteBreedUri = BreedIdRequiredUri;

export const breedEndpoints = {
  addBreed: '/breed',
  getBreed: (id: number) => `/breeds/${id}`,
  putBreed: (id?: number) => id != null ? `/breeds/${id}` : '/breeds',
  patchBreed: (id?: number) => id != null ? `/breeds/${id}` : '/breeds',
  deleetBreed: (id: number) => `/breeds/${id}`,
} as const;

