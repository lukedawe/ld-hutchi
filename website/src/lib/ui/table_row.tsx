import { Table, TextInput, Button, Group, Grid } from "@mantine/core";
import type { CategoryResponse } from "../dtos/responses";
import { cloneElement, useState } from "react";
import { API_BASE_URL } from "../../App";
import BreedList from "./breed_list";

export function CategoryRow(
    { category, deleteCategory, setError, setMessage }: {
        category: CategoryResponse,
        deleteCategory: (breedId: number) => void,
        setError: (error: string) => void,
        setMessage: (message: string) => void,
    }) {
    // add state for the editable category name and breed values (controlled)
    const [categoryName, setCategoryName] = useState(category.name);
    const [breedNames, setBreedNames] = useState<string[]>(
        () => category.breeds.map((b) => b.name)
    );
    const [submitting, setSubmitting] = useState(false);

    // dirty if category name or any breed name differs from original
    const isDirty =
        categoryName !== category.name ||
        breedNames.some((n, i) => n !== (category.breeds[i]?.name ?? ""));

    const submitChanges = async () => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/category/${category.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: categoryName, breeds: breedNames.map((name) => ({ name })) }),
            });

            if (!res.ok) {
                // try to read error body if present
                let errText = '';
                try { errText = await res.text(); } catch (_) { /* ignore */ }
                throw new Error(`Request failed ${res.status}${errText ? `: ${errText}` : ''}`);
            }

            const responseJson = await res.json().catch(() => null);

            if (!responseJson || typeof responseJson !== 'object') {
                throw new Error('Invalid response from server');
            }

            // safe check for expected fields before casting
            if ('name' in responseJson && 'breeds' in responseJson && Array.isArray((responseJson as any).breeds)) {
                const categoryResponse = responseJson as CategoryResponse;
                setCategoryName(categoryResponse.name);
                setBreedNames(categoryResponse.breeds.map((breed) => breed.name));
                // notify success
                setMessage('Category saved');
            } else if ('message' in responseJson && typeof (responseJson as any).message === 'string') {
                throw new Error((responseJson as any).message);
            } else {
                throw new Error('Unexpected server response');
            }
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
                try { errText = await res.text(); } catch (_) { }
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

    const addBreedField = () => {
        setBreedNames((prev)=>[...prev, ""])
    }

    return (
        <Table.Tr key={category.id}>
            <Table.Td>
                <Group align="center">
                    <TextInput
                        value={categoryName}
                        onChange={(e) => setCategoryName(e.currentTarget.value)}
                    />
                </Group>
            </Table.Td>
            <Table.Td>
                <BreedList categoryId={category.id} breedNames={breedNames} setBreedNames={setBreedNames} addBreed={addBreedField}></BreedList>
            </Table.Td>
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
        </Table.Tr>
    );
}