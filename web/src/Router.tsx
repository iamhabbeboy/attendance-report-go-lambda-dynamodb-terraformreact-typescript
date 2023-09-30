import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { HomePage } from './pages/Home.page';
import InsightPage from './pages/Insight.page';

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: '/insight/:id',
    element: <InsightPage />,
  },
]);

export function Router() {
  return <RouterProvider router={router} />;
}
