import { Button, Card, Grid, TextInput } from "@mantine/core";
import type { BreedResponse } from "../dtos/responses";


interface BreedListProps {
    categoryId: number | undefined,
    breeds: Array<BreedResponse>,
    setBreedName: (id: number, name: string) => void,
    addBreed: () => void,
    deleteBreed: (id: number) => void,
    isLoading: boolean,
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

    const breedsList = breeds.map((breed) => (
        <Grid mt="md">
            <BreedCard
                key={`${breed.id}`}
                value={breed.name}
                onChange={(v) => setBreedName(breed.id, v)}
            />
            <Button color="red" size="xs" onClick={()=>deleteBreed(breed.id)}>
                Remove
            </Button>
        </Grid>
    ));

    return (
        <>
            {breedsList}
            <Button onClick={addBreed} variant="filled" radius="lg" size="xs">Add Breed</Button>
        </>
    )
}