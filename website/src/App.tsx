import { useState, useEffect, createContext } from 'react'
import './App.css'
import './lib/dtos/responses'
import type { CategoryResponse } from './lib/dtos/responses'
import { createTheme, Group, MantineProvider, Table, type NotificationProps } from '@mantine/core'
import { notifications, Notifications } from '@mantine/notifications'
import '@mantine/core/styles.css';
import '@mantine/notifications/styles.css';
import { useFetch } from '@mantine/hooks'
import { CategoryRow } from './lib/ui/CategoryRow'

const theme = createTheme({
  /** Put your mantine theme override here */
});

export const API_BASE_URL = 'http://localhost:5173/services/v1'

function App() {
  const [currentPage, setCurrentPage] = useState(1);
  const { data, loading, error, refetch, abort } = useFetch<CategoryResponse[]>(
    `${API_BASE_URL}/breeds/categories/` + currentPage + "/" + 100,
    { autoInvoke: true }
  )

  const [categories, setCategories] = useState(
    new Map(data?.map(
      item => [item.id, item] as const
    )
    )
  )

  useEffect(() =>
    setCategories(
      new Map(data?.map(
        item => [item.id, item] as const
      )
      )
    ),
    [data]
  )

  // TODO: Be able to display errors and info via context instead of via callbacks 

  const setMessage = (message: string) => notifications.show({ title: "Success!", message: message });
  const setError = (errorMessage: string) => notifications.show({ title: "Error occurred", message: errorMessage });

  if (error) {
    console.log("Something went wrong with fetching from the server.");
  }

  // const deleteCategorySuccess = (transform : (prev: CategoryResponse[])=>CategoryResponse[]) => setCategories((prev) => (transform(prev)));
  const deleteCategory = (id: number) => {
    setCategories(prev => {
      prev.delete(id)
      // Trigger re-rendering!
      return new Map(prev)
    }
    )
  }

  const setCategoryState = (category: CategoryResponse) => {
    setCategories((prev) => {
      prev.set(category.id, category);
      return new Map(prev);
    })
  }

  const rows = Array.from(categories).map(([_, category]) => (
    <CategoryRow category={category} deleteCategorySuccessful={deleteCategory} setCategoryState={setCategoryState} setError={setError} setMessage={setMessage} />
  ));

  useEffect(() => {
    refetch();
  }, [currentPage]);

  return (
    <MantineProvider theme={theme}>
      <Notifications />
      {data ?
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
            {/* // TODO: Remove this.
            <Button onClick={async () => {
              // create a new category with an empty name, then add to map when server responds
              try {
                const res = await fetch(`${API_BASE_URL}/category`, { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ name: "" }) });
                if (!res.ok) throw new Error(`Create category failed ${res.status}`);
                const payload = await res.json();
                const created: CategoryResponse = { id: payload.id ?? Date.now(), name: String(payload.name ?? ""), breeds: payload.breeds ?? [] };
                setCategories((prev) => new Map(prev).set(created.id, created));
                setMessage('Category created');
              } catch (err) {
                const message = err instanceof Error ? err.message : String(err);
                setError(message);
              }
            }}>
              Add Category
            </Button> */}
          </Group>
        </>
        : null}
    </MantineProvider>
  )
}

export default App
