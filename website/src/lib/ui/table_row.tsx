import { Card, Table, Text } from "@mantine/core";
import type { CategoryResponse } from "../dtos";

export function row(category: CategoryResponse) {
    const breedsList = category.breeds.map((breed) => {
        return (<Card shadow="sm" padding="lg">
            <Card.Section>
                <Text>
                    {breed.name}
                </Text>
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