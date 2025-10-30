import { useState, useEffect } from 'react'
import './App.css'
import './lib/dtos'
import type { CategoryResponse } from './lib/dtos'
import { createTheme, MantineProvider, Table } from '@mantine/core'
import { CategoryRow } from './lib/ui/table_row'
import '@mantine/core/styles.css';


const theme = createTheme({
  /** Put your mantine theme override here */
});

const API_BASE_URL = 'http://localhost:5173/services/v1'

function App() {
   const [categories, setCategories] = useState<CategoryResponse[] | null>(
     [{ id: 1, name: "hello", breeds: [] }, { id: 2, name: "hello", breeds: [{ name: "Doggy", id: 4 }, { name: "Brother", id: 4 }] }]
 
   )
   const [loading, setLoading] = useState(false);
   const [error, setError] = useState<string | null>(null);
   // id of expanded category (if you need to show/hide details)
   const [expandedCategoryId, setExpandedCategoryId] = useState<number | null>(null);
   const [currentPage, setCurrentPage] = useState(1);

   const fetchCategories = async () => {
     setLoading(true);
     setError(null);

     try {
       const endpoint = `${API_BASE_URL}/breeds/categories/` + currentPage + "/" + 100;
       const response = await fetch(endpoint)

       if (!response.ok) {
         throw new Error(`HTTP error; status: ${response.status}`);
       }

       const data = await response.json();
       console.log(data)
       setCategories(data);
     }
     catch (err) {
       console.error('Error fetching all categories', err);
       setError("Endpoint not reachable.");
       setCategories(null);
     }
     finally {
       setLoading(false);
     }
   }

   useEffect(() => {
     fetchCategories();
   }, [currentPage]);

   const rows = categories?.map((category) => (
     <CategoryRow key={category.id} category={category} />
   ));

   return (
     <MantineProvider theme={theme}>
       <Table striped highlightOnHover withTableBorder>
         <Table.Thead>
           <Table.Tr id='heading'>
             <Table.Th>
               Name
             </Table.Th>
             <Table.Th>
               Breeds
             </Table.Th>
           </Table.Tr>
         </Table.Thead>
         <Table.Tbody>{rows}</Table.Tbody>
       </Table>
     </MantineProvider>
   )
 }

 export default App
