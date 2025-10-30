import { Table } from "@mantine/core";
import type { CategoryResponse } from "../dtos/responses";
import { CategoryRow } from "./table_row";


export default function CategoriesTable({categories}: {categories: CategoryResponse[] | null}) {
    const rows = categories?.map((category) => (
        <CategoryRow key={category.id} category={category} />
    ));

    return (
        <Table striped highlightOnHover withTableBorder>
            <Table.Thead>
                <Table.Tr id='heading'>
                    <Table.Th>
                        Name
                    </Table.Th>
                    <Table.Th>
                        Breeds
                    </Table.Th>
                    <Table.Th>
                        Actions
                    </Table.Th>
                </Table.Tr>
            </Table.Thead>
            <Table.Tbody>{rows}</Table.Tbody>
        </Table>
    );
}     
