import { Table, TextInput, Button, Group, Grid } from "@mantine/core";
import { BreedResponseZod, CategoryResponseZod, ErrorResponseZod, type BreedResponse, type CategoryResponse, type ErrorResponse } from "../dtos/responses";
import { useState, useRef, useEffect } from "react";
import { API_BASE_URL } from "../../App";
import BreedList from "./breed_list";
import type { AddBreed } from "../dtos/requests/breed";

interface CategoryRowProps {
    category: CategoryResponse;
    deleteCategorySuccessful: (id: number) => void;
    setCategoryState: (change: CategoryResponse) => void;
    setError: (error: string) => void;
    setMessage: (message: string) => void;
}

export function CategoryRow({ category, deleteCategorySuccessful, setError, setMessage, setCategoryState }: CategoryRowProps) {
    // local editable state; keep original for reset
    const [submitting, setSubmitting] = useState(false);

    const deleteCategoryRequest = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${category.id}`, { method: 'DELETE' });
            if (!res.ok) {
                const errorMessage = ErrorResponseZod.safeParse(await res.json())
                if (errorMessage.success) {
                    setError((errorMessage.data as ErrorResponse).message)
                    return;
                }
                else setError("Unexpected result");
            }
            deleteCategorySuccessful(category.id);
            setMessage('Category deleted');
        } catch (err) {
            console.error('delete category error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    };

    // Add a new breed by POSTing to the API and appending the returned breed
    const addBreedRequest = async (request: AddBreed) => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/breed`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(request),
            });
            if (!res.ok) {
                const errorMessage = ErrorResponseZod.safeParse(await res.json())
                if (errorMessage.success) {
                    setError((errorMessage.data as ErrorResponse).message);
                    return;
                }
                else setError("Unexpected result");
            }
            const breedResponseJson = await res.json();
            const breedResponse = BreedResponseZod.parse(breedResponseJson);
            setCategoryState({ ...category, breeds: [...category.breeds, breedResponse as BreedResponse] });
            setMessage('Breed added');
        } catch (err) {
            console.error('addBreed error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    };

    const removeBreedRequest = async (id: number) => {
        setSubmitting(true);

        try {
            const res = await fetch(`${API_BASE_URL}/breed`, {
                method: 'DELETE',
            });
            if (!res.ok) {
                const errorMessage = ErrorResponseZod.safeParse(await res.json())
                if (errorMessage.success) {
                    setError((errorMessage.data as ErrorResponse).message)
                    return;
                }
                else setError("Unexpected result");
            }
            setCategoryState({ ...category, breeds: category.breeds.filter((b) => b.id !== id) });
        } catch (err) {
            console.error('addBreed error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
        
    };

    const setBreedName = (id: number, name: string) => {
        // TODO: Add sending request here.
        setCategoryState({ ...category, breeds: category.breeds.map((breed) => breed.id !== id ? breed : { id: breed.id, name }) });
    };

    const updateCategoryName = (newName: string) => {
        // Run query 

        setCategoryState({ ...category, name: newName });
    };

    const resetChanges = () => {
        // setCategoryState(originalRef.current);
    };

    // Reset text fields when new category state is pushed.
    useEffect(
        () => {
            setCategoryNameTextFieldValue(category.name);
        },
        [category]
    )

    const [categoryNameTextFieldValue, setCategoryNameTextFieldValue] = useState(category.name)

    return (
        <Table.Tr key={category.id}>
            <Table.Td>
                <Group align="center">
                    <TextInput
                        value={categoryNameTextFieldValue}
                        onChange={(e) => setCategoryNameTextFieldValue(e.currentTarget.value)}
                    />
                    {
                        categoryNameTextFieldValue !== category.name &&
                        <Button onClick={() => updateCategoryName(categoryNameTextFieldValue)}>
                            Submit changes
                        </Button>
                    }
                </Group>
            </Table.Td>
            <Table.Td>
                <BreedList
                    categoryId={category.id}
                    breeds={category.breeds}
                    setBreedName={setBreedName}
                    addBreed={(request) => addBreedRequest(request)}
                    deleteBreed={removeBreedRequest}
                    isLoading={submitting}
                />
            </Table.Td>
            <Table.Td>
                <Grid w={400}>
                    <Grid.Col>
                        <Button color="red" size="xs" onClick={deleteCategoryRequest} loading={submitting}>
                            Delete
                        </Button>
                    </Grid.Col>
                    {/* <Grid.Col hidden={!isDirty}>
                        <Group>
                            <Button size="xs" onClick={resetChanges} disabled={submitting}>
                                Reset
                            </Button>
                            <Button size="xs" onClick={submitChangesToCategoryRequest} loading={submitting}>
                                Save
                            </Button>
                        </Group>
                    </Grid.Col> */}
                </Grid>
            </Table.Td>
        </Table.Tr>
    );
}