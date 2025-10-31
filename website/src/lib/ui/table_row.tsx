import { Table, TextInput, Button, Group, Grid } from "@mantine/core";
import { jsonToCategoryResponse, type BreedResponse, type CategoryResponse } from "../dtos/responses";
import { cloneElement, useState } from "react";
import { API_BASE_URL } from "../../App";
import BreedList from "./breed_list";

interface CategoryRow {
    category: CategoryResponse,
    deleteCategory: (CategoryId: number) => void,
    deleteBreed: (breedId: number) => void,
    addBreed: () => void,
    setError: (error: string) => void,
    setMessage: (message: string) => void,
}

export function CategoryRow(
    { category, deleteCategory, setError, setMessage }: CategoryRow) {
    // add state for the editable category name and breed values (controlled).
    const [categoryState, setCategoryState] = useState(category);
    const [submitting, setSubmitting] = useState(false);

    // dirty if category name or any breed name differs from original.
    const isDirty =
        categoryState.name !== category.name ||
        categoryState.breeds.some((breed, i) => breed.name !== (category.breeds[i]?.name ?? ""));

    const submitChanges = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${category.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    id: categoryState.id,
                    name: categoryState.name,
                    breeds: categoryState.breeds,
                }
                ),
            });

            if (!res.ok) {
                // try to read error body if present
                let errText = '';
                try { errText = (await res.json()).message; } catch (_) { /* ignore */ }
                throw new Error(`Request failed ${res.status}${errText ? `: ${errText}` : ''}`);
            }

            const responseJson = await res.json().catch(() => null);
            // Can throw casting errors.
            const categoryResponse = jsonToCategoryResponse(responseJson);
            setCategoryState(categoryResponse);
            setMessage('Category saved');

        } catch (err) {
            console.error('submitChanges error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    }

    const deleteColumn = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${category.id}`, { method: 'DELETE' });
            if (!res.ok) {
                let errText = '';
                try { errText = (await res.json()).message; } catch (_) { }
                throw new Error(`Delete failed ${res.status}${errText ? `: ${errText}` : ''}`);
            }
            deleteCategory(category.id);
            setMessage('Category deleted');
        } catch (err) {
            console.error('deleteColumn error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    }

    const addBreed = () => {
        setCategoryState(currentVal => {
            currentVal.breeds = [...currentVal.breeds, { id: 0, name: "" }];
            return currentVal;
        })
    }

    // TODO: Submit a network request before settng the state.
    const removeBreed = (id: number) => {
        setCategoryState(currentVal => {
            currentVal.breeds = currentVal.breeds.filter(element => element.id !== id)
            return currentVal;
        });
    }

    const setBreedName = (id: number, name: string) => {
        setCategoryState(currentVal => {
            currentVal.breeds = currentVal.breeds.map(breed => breed.id !== id ? breed : { id: breed.id, name: name })
            return currentVal;
        });
    }

    const setCategoryName = (newName: string) => {
        setCategoryState(currentVal => {
            currentVal.name = newName;
            return currentVal;
        }
        )
    }

    return (
        <Table.Tr key={category.id}>
            <Table.Td>
                <Group align="center">
                    <TextInput
                        value={categoryState.name}
                        onChange={(e) => setCategoryName(e.currentTarget.value)}
                    />
                </Group>
            </Table.Td>
            <Table.Td>
                <BreedList
                    categoryId={category.id}
                    breeds={category.breeds}
                    setBreedName={setBreedName}
                    addBreed={addBreed}
                    deleteBreed={removeBreed}
                    isLoading={submitting}
                />
            </Table.Td>
            <Table.Td>
                <Grid w={400}>
                    <Grid.Col>
                        <Button color="red" size="xs" onClick={deleteColumn} loading={submitting}>
                            Delete
                        </Button>
                    </Grid.Col>
                    <Grid.Col hidden={!isDirty}>
                        <Button size="xs" onClick={submitChanges} loading={submitting}>
                            Save
                        </Button>
                    </Grid.Col>
                </Grid>
            </Table.Td>
        </Table.Tr>
    );
}