import { useState, useEffect } from 'react'
import './App.css'
import './lib/dtos/responses'
import type { CategoryResponse } from './lib/dtos/responses'
import { createTheme, MantineProvider, Table, type NotificationProps } from '@mantine/core'
import {notifications, Notifications} from '@mantine/notifications'
import '@mantine/core/styles.css';
import '@mantine/notifications/styles.css';
import CategoriesTable from './lib/ui/categories_table'
import { useFetch } from '@mantine/hooks'

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

  if (error) {
    console.log("Something went wrong with fetching from the server.");
  }

  useEffect(() => {
    refetch();
  }, [currentPage]);

  return (
    <MantineProvider theme={theme}>
      <Notifications />
      {data ? 
      <CategoriesTable
        categories={data}
        setMessage={(message) => notifications.show({ title:"Success!", message: message })}
        setError={(errorMessage) => notifications.show({title:"Error occurred", message: errorMessage})}
      /> : null }
    </MantineProvider>
  )
}

export default App
