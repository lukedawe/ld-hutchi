import { Badge, Group, Table, Text } from "@mantine/core";
import type { BreedResponse, CategoryResponse } from "../dtos";

function row(category:CategoryResponse) {
    return (
        <Table.Tr>
            <Table.Td>
                <Badge color="blue" variant="light" size="lg">{category.id}</Badge>
            </Table.Td>
            <Table.Td>
                <Group gap="xs">
                    <Text fw={500}>{category.name}</Text>
                </Group>
            </Table.Td>
        </Table.Tr>
    );
}