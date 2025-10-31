import { Card, Table, TextInput, Button, Group, Grid } from "@mantine/core";
import type { CategoryResponse } from "../dtos/responses";
import { useState } from "react";
import { API_BASE_URL } from "../../App";

export function CategoryRow(
    { category, deleteCategory, setError }: {
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

    function BreedCard({ value, onChange }: { value: string; onChange: (v: string) => void }) {
        return (
            <Card shadow="sm" padding="lg">
                <Card.Section>
                    <TextInput value={value} onChange={(event) => onChange(event.currentTarget.value)} />
                </Card.Section>
            </Card>
        );
    }

    const breedsList = breedNames.map((bName, idx) => (
        <BreedCard
            key={`${category.breeds[idx]?.id ?? 'breed'}-${idx}`}
            value={bName}
            onChange={(v) => setBreedNames((prev) => {
                const next = [...prev];
                next[idx] = v;
                return next;
            })}
        />
    ));

    // dirty if category name or any breed name differs from original
    const isDirty =
        categoryName !== category.name ||
        breedNames.some((n, i) => n !== (category.breeds[i]?.name ?? ""));

    const submitChanges = async () => {
        setSubmitting(true);
        fetch(`${API_BASE_URL}/category/${category.id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                name: categoryName,
                breeds: breedNames.map((name)=>({name: name})),
            })
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response;
            })
            .then((response) => response.json())
            .then((responseJson) => {
                const categoryResponse = responseJson as CategoryResponse
                setCategoryName(categoryResponse.name);
                setBreedNames(categoryResponse.breeds.map((breed) => breed.name))
            }
            )
            .catch(error => {
                console.log("there has been an error.", error);
                const message = (error && (error as any).message) ? (error as any).message : String(error);
                setError(message);
            }
            )
            .finally(() => setSubmitting(false))
    }

    const deleteColumn = async () => {
        setSubmitting(true);
        fetch(`${API_BASE_URL}/category/${category.id}`, {
            method: 'DELETE'
        })
            .then((response) => response.ok ? deleteCategory(category.id) : null)
            .catch(error => {
                console.log(`there has been an error:`, error)
                const message = (error && (error as any).message) ? (error as any).message : String(error);
                setError(message);
            })
            .finally(() => setSubmitting(false))
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
            <Table.Td>{breedsList}</Table.Td>
            <Grid w={400}>
                <Grid.Col>
                    <Button size="xs" onClick={deleteColumn} loading={submitting}>
                        Delete
                    </Button>
                </Grid.Col>
                {isDirty && (
                    <Grid.Col>
                        <Button size="xs" onClick={submitChanges} loading={submitting}>
                            Save
                        </Button>
                    </Grid.Col>
                )
                }
            </Grid>
        </Table.Tr>
    );
}