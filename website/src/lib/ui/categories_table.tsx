import { Table } from "@mantine/core";
import type { CategoryResponse } from "../dtos/responses";
import { CategoryRow } from "./table_row";
import { useState, useEffect } from "react";


export default function CategoriesTable(
    { categories, setError, setMessage }: {
        categories: CategoryResponse[],
        setError: (error: string) => void,
        setMessage: (message: string) => void
    }) {

    // initialize a Map keyed by category id from the incoming categories array
    const [categoryMap, setCategoryMap] = useState<Map<number, CategoryResponse>>(() =>
        new Map(categories?.map((c) => [c.id, c] as [number, CategoryResponse]) ?? [])
    );

    // keep local map in sync if the categories prop changes
    useEffect(() => {
        setCategoryMap(new Map(categories?.map((c) => [c.id, c] as [number, CategoryResponse]) ?? []));
    }, [categories]);

    const deleteCategory = (id: number) => setCategoryMap((prev) => {
        const clone = new Map(prev);
        clone.delete(id);
        return clone;
    });

    const rows = Array.from(categoryMap.values()).map((category) => (
        <CategoryRow key={category.id} category={category} deleteCategory={deleteCategory} setError={setError} setMessage={setMessage} />
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
