import { Table, TextInput, Button, Group, Grid } from "@mantine/core";
import { ErrorResponseZod, type CategoryResponse, type ErrorResponse } from "../dtos/responses";
import { useState, useRef, useEffect } from "react";
import { API_BASE_URL } from "../../App";
import BreedList from "./BreedList";

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
    const [categoryNameTextFieldValue, setCategoryNameTextFieldValue] = useState(category.name)

    const deleteCategoryRequest = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${category.id}`, { method: 'DELETE' });
            if (!res.ok) {
                const errorMessage = ErrorResponseZod.safeParse(await res.json())
                if (errorMessage.success) {
                    setError((errorMessage.data as ErrorResponse).message)
                }
                else setError("Unexpected result");

                return;
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

    const updateCategoryNameRequest = async (newName: string) => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${category.id}`,
                {
                    method: 'UPDATE',
                    body: JSON.stringify(
                        {
                            name: newName,
                        }
                    )
                });
            if (!res.ok) {
                const errorMessage = ErrorResponseZod.safeParse(await res.json())
                if (errorMessage.success) {
                    setError((errorMessage.data as ErrorResponse).message)
                }
                else setError("Unexpected result");

                return;
            }
            setCategoryState({ ...category, name: newName })
            setMessage('Category deleted');
        } catch (err) {
            console.error('delete category error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    };

    // Reset text fields when new category state is pushed.
    useEffect(
        () => {
            setCategoryNameTextFieldValue(category.name);
        },
        [category]
    )

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
                        <>
                            <Button onClick={() => updateCategoryNameRequest(categoryNameTextFieldValue)}>
                                Submit changes
                            </Button>
                            <Button onClick={() => setCategoryNameTextFieldValue(category.name)}>
                                Reset
                            </Button>
                        </>
                    }
                </Group>
            </Table.Td>
            <Table.Td>
                <BreedList
                    categoryId={category.id}
                    breedList={category.breeds}
                    changeBreedList={(newBreeds) => setCategoryState({ ...category, breeds: newBreeds })}
                    submitting={submitting}
                    setSubmitting={setSubmitting}
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