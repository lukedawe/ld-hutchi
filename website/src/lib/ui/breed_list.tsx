import { Button, Card, Grid, TextInput, Group } from "@mantine/core";
import type { BreedResponse } from "../dtos/responses";

interface BreedListProps {
    categoryId: number | undefined;
    breeds: Array<BreedResponse>;
    setBreedName: (id: number, name: string) => void;
    addBreed: () => Promise<void> | void;
    deleteBreed: (id: number) => void;
    isLoading: boolean;
}

export default function BreedList({ categoryId, breeds, setBreedName, addBreed, deleteBreed, isLoading }: BreedListProps) {
    function BreedCard({ value, onChange }: { value: string; onChange: (v: string) => void }) {
        return (
            <Card shadow="sm" padding="lg">
                <Card.Section>
                    <TextInput disabled={isLoading} value={value} onChange={(event) => onChange(event.currentTarget.value)} />
                </Card.Section>
            </Card>
        );
    }

    return (
        <>
            {breeds.map((breed) => (
                <Grid key={breed.id} mt="md" align="center">
                    <Grid.Col span={8}>
                        <BreedCard
                            value={breed.name}
                            onChange={(v) => setBreedName(breed.id, v)}
                        />
                    </Grid.Col>
                    <Grid.Col span={4}>
                        <Group>
                            <Button color="red" size="xs" onClick={() => deleteBreed(breed.id)}>
                                Remove
                            </Button>
                        </Group>
                    </Grid.Col>
                </Grid>
            ))}

            <Group mt="md">
                <Button onClick={() => void addBreed()} variant="filled" radius="lg" size="xs" disabled={isLoading}>Add Breed</Button>
            </Group>
        </>
    );
}