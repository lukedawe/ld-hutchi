import { Button, Card, CloseButton, TextInput } from "@mantine/core";
import type { Dispatch, SetStateAction } from "react";


export default function BreedList({ categoryId, breedNames, setBreedNames, addBreed }: { categoryId: number | undefined, breedNames: Array<string>, setBreedNames: Dispatch<SetStateAction<string[]>>, addBreed: () => void}) {
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
            key={`${categoryId ?? 'breed'}-${idx}`}
            value={bName}
            onChange={(v) => setBreedNames((prev) => {
                const next = [...prev];
                next[idx] = v;
                return next;
            })}
        />
    ));

    return (
        <>
            {breedsList}
            <Button onClick={addBreed} variant="filled" radius="lg" size="xs">Add Breed</Button>
        </>
    )
}