import { Card, Table, TextInput } from "@mantine/core";
import type { CategoryResponse } from "../dtos";
import { useState } from "react";

export function row(category: CategoryResponse) {
    const breedsList = category.breeds.map((breed) => {
        const [value, setValue] = useState(breed.name)
        return (<Card shadow="sm" padding="lg">
            <Card.Section>
                <TextInput
                    value={value}
                    onChange={(event) => setValue(event.currentTarget.value)}
                />
            </Card.Section>
        </Card>)
    });

    return (
        <Table.Tr key={category.id}>
            <Table.Td>{category.name}</Table.Td>
            <Table.Td>
                {breedsList}
            </Table.Td>
        </Table.Tr>
    );
}