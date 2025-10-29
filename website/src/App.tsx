import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import './lib/dtos'
import type { CategoryResponse } from './lib/dtos'

const API_BASE_URL = 'http://localhost:8081/v1'

function App() {
  const [count, setCount] = useState(0)
  const [categories, setCategories] = useState<CategoryResponse[] | null>([])
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  // A function to show and hide the breeds.
  const [openCategory, setOpenCategory] = useState<number | null>(null);
  const [currentPage, setCurrentPage] = useState(1);

  const fetchCategories = async () => {
    setLoading(true);
    setError(null);

    try {
      const endpoint = `${API_BASE_URL}/breeds/categories/` + currentPage;
      const response = await fetch(endpoint)

      if (!response.ok) {
        throw new Error(`HTTP error; status: ${response.status}`);
      }

      const data = await response.json();
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
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
