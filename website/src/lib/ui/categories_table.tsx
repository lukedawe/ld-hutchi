import { Table } from "@mantine/core";
import type { CategoryResponse } from "../dtos/responses";
import { CategoryRow } from "./table_row";
import { useState, useEffect } from "react";
import { Button, Group } from "@mantine/core";
import { API_BASE_URL } from "../../App";


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
        <CategoryRow key={category.id} category={category} deleteCategorySuccessful={deleteCategory} setError={setError} setMessage={setMessage} />
    ));

    return (
        <>
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
            <Group mt="md">
                <Button onClick={async () => {
                    // create a new category with an empty name, then add to map when server responds
                    try {
                        const res = await fetch(`${API_BASE_URL}/category`, { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ name: "" }) });
                        if (!res.ok) throw new Error(`Create category failed ${res.status}`);
                        const payload = await res.json();
                        const created: CategoryResponse = { id: payload.id ?? Date.now(), name: String(payload.name ?? ""), breeds: payload.breeds ?? [] };
                        setCategoryMap((prev) => new Map(prev).set(created.id, created));
                        setMessage('Category created');
                    } catch (err) {
                        const message = err instanceof Error ? err.message : String(err);
                        setError(message);
                    }
                }}>
                    Add Category
                </Button>
            </Group>
        </>
    );
}
