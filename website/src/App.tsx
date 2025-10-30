import { useState, useEffect } from 'react'
import './App.css'
import './lib/dtos/responses'
import type { CategoryResponse } from './lib/dtos/responses'
import { createTheme, MantineProvider, Table } from '@mantine/core'
import { CategoryRow } from './lib/ui/table_row'
import '@mantine/core/styles.css';
import CategoriesTable from './lib/ui/categories_table'
import { useFetch } from '@mantine/hooks'

const theme = createTheme({
  /** Put your mantine theme override here */
});

const API_BASE_URL = 'http://localhost:5173/services/v1'

function App() {
  const [currentPage, setCurrentPage] = useState(1);
  const { data, loading, error, refetch, abort } = useFetch<CategoryResponse[]>(
    `${API_BASE_URL}/breeds/categories/` + currentPage + "/" + 100,
    {autoInvoke: true}
  )

  if (error) {
    console.log("Something went wrong with fetching from the server.");
  }

  useEffect(() => {
    refetch();
  }, [currentPage]);

  return (
    <MantineProvider theme={theme}>
      <CategoriesTable categories={data} />
    </MantineProvider>
  )
}

export default App
