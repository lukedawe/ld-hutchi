import { RequestInfo } from "./requestInfo";

export interface CategoryNameRequiredJson {
  name: string;
}

export interface CategoryNameRequiredUri {
  name: string;
}

export interface CategoryIdRequiredUri {
  id: number;
}

export interface CategoryIdOptionalUri {
  id?: number;
}

/** Breed shape used inside category requests */
export interface Breed {
  name: string;
}

export interface BreedArrayRequired {
  breeds: Breed[];
}

/** AddCategoryJson has both name and breeds */
export interface AddCategoryJson {
  name: string;
  breeds: Breed[];
}

/** Wrapper for multiple categories */
export interface AddCategories {
  categories: AddCategoryJson[];
}

export interface Paginated {
  page: number;
  per_page: number;
}

/** Classes that implement the request shapes and provide toRequestInfo() */
export class GetCategoriesToBreeds implements Paginated {
  page: number;
  per_page: number;

  constructor(page: number, per_page: number) {
    this.page = page;
    this.per_page = per_page;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'GetCategoriesToBreeds', page: this.page, per_page: this.per_page } });
  }
}

export class GetCategory implements CategoryIdRequiredUri {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'GetCategory', id: this.id } });
  }
}

export class GetCategoryToBreeds implements CategoryIdRequiredUri {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'GetCategoryToBreeds', id: this.id } });
  }
}

export class PutCategoryUri implements CategoryIdOptionalUri {
  id?: number;

  constructor(id?: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'PutCategoryUri', id: this.id } });
  }
}

export class PatchCategoryUri implements CategoryIdRequiredUri {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'PatchCategoryUri', id: this.id } });
  }
}

export class PatchCategoryBody implements CategoryNameRequiredJson {
  name: string;

  constructor(name: string) {
    this.name = name;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'PatchCategoryBody', name: this.name } });
  }
}

export class DeleteCategoryUri implements CategoryIdRequiredUri {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  toRequestInfo() {
    return new RequestInfo({ body: { class: 'DeleteCategoryUri', id: this.id } });
  }
}

export const categoryEndpoints = {
  addCategory: '/categories',
  addCategories: '/categories/bulk',
  getCategories: (page: Paginated) => {
    let url = '/categories';
    const params: string[] = [];
   params.push(`page=${page.page}`);
    params.push(`per_page=${page.per_page}`);
    if (params.length) url += `?${params.join('&')}`;
    return url;
  },
  getCategory: (id?: number) => (typeof id === 'number' ? `/categories/${id}` : '/categories'),
  getCategoryToBreeds: (id: number) => `/categories/${id}/breeds`,
  putCategory: (id?: number) => (typeof id === 'number' ? `/categories/${id}` : '/categories'),
  patchCategory: (id: number) => `/categories/${id}`,
  deleteCategory: (id: number) => `/categories/${id}`,
} as const;