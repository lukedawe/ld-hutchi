import { Card, Table, TextInput } from "@mantine/core";
import type { CategoryResponse } from "../dtos";
import { useState } from "react";

export function CategoryRow({ category }: { category: CategoryResponse }) {
    // Component for an editable breed card
    function BreedCard({ initialName }: { initialName: string }) {
        const [value, setValue] = useState(initialName);
        return (
            <Card shadow="sm" padding="lg">
                <Card.Section>
                    <TextInput
                        value={value}
                        onChange={(event) => setValue(event.currentTarget.value)}
                    />
                </Card.Section>
            </Card>
        );
    }

    const breedsList = category.breeds.map((breed, idx) => (
        <BreedCard
            key={`${breed.id ?? 'breed'}-${idx}`}
            initialName={breed.name}
        />
    ));

    return (
        <Table.Tr key={category.id}>
            <Table.Td>{category.name}</Table.Td>
            <Table.Td>
                {breedsList}
            </Table.Td>
        </Table.Tr>
    );
}