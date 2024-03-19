// @ts-ignore
import React from 'react';
import { render } from '@testing-library/react';
import Table from './Table';

describe('Table Component', () => {
  const testData = [
    { id: 1, name: 'Alice', age: 25 },
    { id: 2, name: 'Bob', age: 30 },
    { id: 3, name: 'Charlie', age: 35 },
  ];

  test('renders correct rows and columns based on data', () => {
    const { getAllByRole } = render(<Table data={testData} />);

    // Check the number of rows
    const rows = getAllByRole('row');
    expect(rows.length).toBe(testData.length + 1); 

    // Check the number of columns in the header row
    const headerRow = rows[0];
    const headerColumns = headerRow.querySelectorAll('th');
    expect(headerColumns.length).toBe(Object.keys(testData[0]).length + 1); 

    // Check the number of columns in each data row
    rows.slice(1).forEach((row) => {
      const dataColumns = row.querySelectorAll('td');
      expect(dataColumns.length).toBe(headerColumns.length);
    });
  });
});
