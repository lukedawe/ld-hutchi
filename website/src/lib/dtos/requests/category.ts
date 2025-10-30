// ...new file...
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

export type GetCategoriesToBreeds = Paginated;
export type GetCategory = CategoryIdRequiredUri;
export type GetCategoryToBreeds = CategoryIdRequiredUri;
export type PutCategoryUri = CategoryIdOptionalUri;
export type PatchCategoryUri = CategoryIdRequiredUri;
export type PatchCategoryBody = CategoryNameRequiredJson;
export type DeleteCategoryUri = CategoryIdRequiredUri;

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