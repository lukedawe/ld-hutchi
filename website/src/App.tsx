import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import './lib/dtos'
import type { CategoryResponse } from './lib/dtos'

const API_BASE_URL = 'http://localhost:8081/v1'

function App() {
  const [count, setCount] = useState(0)
  const [categories, setCategories] = useState<CategoryResponse[] | null>(
    [{ id: 1, name: "hello", breeds: [] }]
  )
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  // A function to show and hide the breeds.
  const [openCategory, setOpenCategory] = useState<number | null>(null);
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
  }, []);

  const toggleBreeds = (categoryId: number) => {
    setOpenCategory(openCategory === categoryId ? null : categoryId);
  }

  return (
    <table>
      <thead>
        {categories?.map(
          (category) =>
            <tr id={category.id.toString()}>
              <th> I am a boss: {category.id}</th>
            </tr>)}
      </thead>
    </table>
  )
}

export default App
