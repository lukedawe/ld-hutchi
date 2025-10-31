// ...new file...
import { RequestInfo } from './requestInfo';

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

/** Classes that implement the request shapes and provide toRequestInfo() */
export class GetBreed implements BreedIdRequiredUri {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'GetBreed', id: this.id } });
  }
}

export class PutBreedUri implements BreedIdOptionalUri {
  id?: number;

  constructor(id?: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'PutBreedUri', id: this.id } });
  }
}

export class PatchBreedUri implements BreedIdOptionalUri {
  id?: number;

  constructor(id?: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'PatchBreedUri', id: this.id } });
  }
}

export class PatchBreedBody implements BreedNameRequiredJson {
  name: string;

  constructor(name: string) {
    this.name = name;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'PatchBreedBody', name: this.name } });
  }
}

export class DeleteBreedUri implements BreedIdRequiredUri {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'DeleteBreedUri', id: this.id } });
  }
}

export const breedEndpoints = {
  addBreed: '/breed',
  getBreed: (id: number) => `/breeds/${id}`,
  putBreed: (id?: number) => id != null ? `/breeds/${id}` : '/breeds',
  patchBreed: (id?: number) => id != null ? `/breeds/${id}` : '/breeds',
  deleetBreed: (id: number) => `/breeds/${id}`,
} as const;

