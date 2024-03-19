// @ts-ignore
import React from 'react';
import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import Layout from './Layout';
import '@testing-library/jest-dom';

//  Mock the Navbar component
jest.mock('../navBar/Navbar', () => () => (
  <nav data-testid='navbar-component'>
    <ul>
      <li>
        <a href='/'>Mocked Navbar Item</a>
      </li>
    </ul>
  </nav>
));

describe('Layout Component', () => {
  test('renders children wrapped in Navbar component', () => {
    // Render the Layout component with a child element
    const { getByTestId } = render(
      <MemoryRouter>
        <Layout>
          <div data-testid='child-element'>Child Element</div>
        </Layout>
      </MemoryRouter>,
    );

    const navbarElement = getByTestId('navbar-component');
    expect(navbarElement).toBeInTheDocument();

    const childElement = getByTestId('child-element');
    expect(childElement).toBeInTheDocument();
    expect(childElement.textContent).toBe('Child Element');
  });
});
