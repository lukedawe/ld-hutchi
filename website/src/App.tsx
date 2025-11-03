import { Button, createTheme, Group, MantineProvider, Modal, Pagination, Table, Text, TextInput } from '@mantine/core'
import '@mantine/core/styles.css'
import { useDisclosure, useFetch } from '@mantine/hooks'
import { notifications, Notifications } from '@mantine/notifications'
import '@mantine/notifications/styles.css'
import { useEffect, useState } from 'react'
import './App.css'
import type { AddCategoryJson } from './lib/dtos/requests/category'
import './lib/dtos/responses'
import { CategoryResponsePaginated, CategoryResponseZod, ErrorResponseZod, type CategoryResponse, type ErrorResponse } from './lib/dtos/responses'
import { CategoryRow } from './lib/ui/CategoryRow'

const theme = createTheme({
  /** Put your mantine theme override here */
});

export const API_BASE_URL = 'http://localhost:5173/services/v1'

function App() {
  const [currentPage, setCurrentPage] = useState(1);
  const { data, loading, error, refetch } = useFetch<CategoryResponsePaginated>(
    `${API_BASE_URL}/breeds/categories/` + currentPage + "/" + 20,
    { autoInvoke: true }
  )

  useEffect(
    () => {
      refetch();
    },
    [currentPage]
  )

  const [categories, setCategories] = useState(
    new Map(data?.categories?.map(
      item => [item.id, item] as const
    )
    )
  )
  // Add category modal.
  const [opened, { open, close }] = useDisclosure(false);

  useEffect(() =>
    setCategories(
      new Map(data?.categories.map(
        item => [item.id, item] as const
      ))
    ),
    [data]
  )

  const setMessage = (message: string) => notifications.show({ title: "Success!", message: message });
  const setError = (errorMessage: string) => notifications.show({ title: "Error occurred", message: errorMessage });

  if (error) {
    console.log("Something went wrong with fetching from the server.");
  }

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

  function AddCategoryModal() {
    const createCategory = async (newCat: AddCategoryJson) => {
      setLoading(true);
      try {
        const res = await fetch(`${API_BASE_URL}/category`,
          {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(newCat)
          });
        if (!res.ok) {
          const errorMessage = ErrorResponseZod.safeParse(await res.json())
          if (errorMessage.success) {
            setError((errorMessage.data as ErrorResponse).message)
          }
          else setError("Unexpected result");

          return;
        }
        const responseJson = await res.json();
        const response = CategoryResponseZod.parse(responseJson) as CategoryResponse;
        console.log(responseJson);
        setCategories((current) => {
          current.set(response.id, response);
          return new Map(current);
        });
        setMessage('Category created');
      } catch (err) {
        console.error('delete category error', err);
        const message = err instanceof Error ? err.message : String(err);
        setError(message);
      } finally {
        setLoading(false);
      }
    }

    const [loading, setLoading] = useState(false)
    const [newCategoryName, setNewCategoryName] = useState("");
    return (
      <Modal opened={opened} onClose={close} title="Add Category">
        <Group>
          <Text>
            New category name:
          </Text>
          <TextInput placeholder='Woof' value={newCategoryName} onChange={(e) => setNewCategoryName(e.target.value)} />
          <Button loading={loading} onClick={() => createCategory({ name: newCategoryName, breeds: [] })}>
            Submit
          </Button>
        </Group>
      </Modal>
    )
  }

  return (
    <MantineProvider theme={theme}>
      <AddCategoryModal></AddCategoryModal>
      <Button onClick={refetch} loading={loading}>
        Reload
      </Button>
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
        </>
        : null}
      <Group mt="md">
        <Button onClick={open}>
          Add Category
        </Button>
      </Group>
      <Pagination total={data?.no_pages ?? 0} onChange={(pageNo) => { setCurrentPage(pageNo) }} />
    </MantineProvider>
  )
}

export default App
