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
