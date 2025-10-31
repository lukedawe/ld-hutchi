import { Table, TextInput, Button, Group, Grid } from "@mantine/core";
import { jsonToCategoryResponse, type BreedResponse, type CategoryResponse } from "../dtos/responses";
import { useState, useRef, useEffect } from "react";
import { API_BASE_URL } from "../../App";
import BreedList from "./breed_list";

interface CategoryRowProps {
    category: CategoryResponse;
    deleteCategory: (CategoryId: number) => void;
    setError: (error: string) => void;
    setMessage: (message: string) => void;
}

export function CategoryRow({ category, deleteCategory, setError, setMessage }: CategoryRowProps) {
    // local editable state; keep original for reset
    const originalRef = useRef<CategoryResponse>(category);
    const [categoryState, setCategoryState] = useState<CategoryResponse>(category);
    // if parent provides a new category object (e.g. refetch), sync local and original
    useEffect(() => {
        originalRef.current = category;
        setCategoryState(category);
    }, [category]);
    const [submitting, setSubmitting] = useState(false);

    // dirty if category name or any breed name differs from original.
    const isDirty =
        categoryState.name !== originalRef.current.name ||
        categoryState.breeds.some((breed, i) => breed.name !== (originalRef.current.breeds[i]?.name ?? ""));

    const submitChanges = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${categoryState.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    id: categoryState.id,
                    name: categoryState.name,
                    breeds: categoryState.breeds,
                }),
            });

            if (!res.ok) {
                let errText = '';
                try { errText = (await res.json()).message; } catch (_) { }
                throw new Error(`Request failed ${res.status}${errText ? `: ${errText}` : ''}`);
            }

            const responseJson = await res.json().catch(() => null);
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
    };

    const deleteColumn = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${categoryState.id}`, { method: 'DELETE' });
            if (!res.ok) {
                let errText = '';
                try { errText = (await res.json()).message; } catch (_) { }
                throw new Error(`Delete failed ${res.status}${errText ? `: ${errText}` : ''}`);
            }
            deleteCategory(categoryState.id);
            setMessage('Category deleted');
        } catch (err) {
            console.error('deleteColumn error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    };

    // Add a new breed by POSTing to the API and appending the returned breed
    const addBreed = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/breed`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ category_id: categoryState.id }),
            });
            if (!res.ok) {
                let errText = '';
                try { errText = (await res.json()).message; } catch (_) { }
                throw new Error(`Add breed failed ${res.status}${errText ? `: ${errText}` : ''}`);
            }
            const payload = await res.json();
            // payload expected to have { id?: number, name: string }
            const newBreed: BreedResponse = { id: payload.id ?? Date.now(), name: String(payload.name ?? "") };
            setCategoryState((cur) => ({ ...cur, breeds: [...cur.breeds, newBreed] }));
            setMessage('Breed added');
        } catch (err) {
            console.error('addBreed error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    };

    const removeBreed = (id: number) => {
        setCategoryState((cur) => ({ ...cur, breeds: cur.breeds.filter((b) => b.id !== id) }));
    };

    const setBreedName = (id: number, name: string) => {
        setCategoryState((cur) => ({ ...cur, breeds: cur.breeds.map((breed) => breed.id !== id ? breed : { id: breed.id, name }) }));
    };

    const setCategoryName = (newName: string) => {
        setCategoryState((cur) => ({ ...cur, name: newName }));
    };

    const resetChanges = () => {
        setCategoryState(originalRef.current);
    };

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
                    categoryId={categoryState.id}
                    breeds={categoryState.breeds}
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
                        <Group>
                            <Button size="xs" onClick={resetChanges} disabled={submitting}>
                                Reset
                            </Button>
                            <Button size="xs" onClick={submitChanges} loading={submitting}>
                                Save
                            </Button>
                        </Group>
                    </Grid.Col>
                </Grid>
            </Table.Td>
        </Table.Tr>
    );
}