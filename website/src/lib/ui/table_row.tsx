import { Card, Table, TextInput, Button, Group, Grid } from "@mantine/core";
import type { CategoryResponse } from "../dtos/responses";
import { useState } from "react";

function SubmitRequest(data: any) {
    // Send post message to api
}

export function CategoryRow({ category }: { category: CategoryResponse }) {
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
        try {
            await fetch('/api/category', {

            });
        } finally {
            setSubmitting(false);
        }
    }

    const deleteColumn = async () => {
        setSubmitting(true);
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