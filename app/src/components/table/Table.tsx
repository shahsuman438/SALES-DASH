import React, { useState, useEffect } from 'react';
import { changeCamelCaseToPascalCaseWithSpaces } from '../../utils/common';

interface Props {
  data: any[];
}

const Table: React.FC<Props> = ({ data }) => {
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(5);
  const [sortConfig, setSortConfig] = useState<{
    key: string;
    direction: 'ascending' | 'descending' | null;
  }>({
    key: '',
    direction: null,
  });

  useEffect(() => {
    setCurrentPage(1);
  }, [itemsPerPage]);

  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  const currentItems = data?.slice(startIndex, endIndex);
  const totalPages = Math.ceil(data?.length / itemsPerPage);

  const handlePageChange = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  const handleItemsPerPageChange = (value: number) => {
    setItemsPerPage(value);
  };

  const sortData = (key: string) => {
    let direction: 'ascending' | 'descending' | null = 'ascending';
    if (sortConfig.key === key && sortConfig.direction === 'ascending') {
      direction = 'descending';
    }
    setSortConfig({ key, direction });
  };

  // Sort data based on the current sort configuration
  let sortedData = [...currentItems];
  if (sortConfig.key !== '') {
    sortedData.sort((a: any, b: any) => {
      const valueA = a[sortConfig.key];
      const valueB = b[sortConfig.key];
      if (typeof valueA === 'string' && typeof valueB === 'string') {
        if (valueA < valueB) {
          return sortConfig.direction === 'ascending' ? -1 : 1;
        }
        if (valueA > valueB) {
          return sortConfig.direction === 'ascending' ? 1 : -1;
        }
        return 0;
      } else if (typeof valueA === 'number' && typeof valueB === 'number') {
        return sortConfig.direction === 'ascending'
          ? valueA - valueB
          : valueB - valueA;
      } else {
        return 0;
      }
    });
  }

  const headers = Object.keys(data[0] || {});

  const itemsPerPageOptions = [5, 10, 20, 40];
  let nextOption = itemsPerPageOptions[itemsPerPageOptions.length - 1] * 2;
  while (nextOption <= data?.length) {
    itemsPerPageOptions.push(nextOption);
    nextOption *= 2;
  }

  return (
    <div className='table-container p-1 pt-4'>
      <table className='table'>
        <thead>
          <tr>
            <th key='sn'>SN</th>
            {headers.map((header, index) => (
              <th key={index} onClick={() => sortData(header)}>
                {changeCamelCaseToPascalCaseWithSpaces(header)}
                {sortConfig.key === header &&
                  (sortConfig.direction === 'ascending' ? ' ▲' : ' ▼')}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {sortedData?.map((item: any, index: number) => (
            <tr key={index}>
              <td>{index + 1}</td>
              {headers.map((header, index) => (
                <td key={index}>{item[header]}</td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>

      <div className='pagination'>
        <div>
          Rows per page:
          <select
            value={itemsPerPage}
            onChange={(e) => handleItemsPerPageChange(parseInt(e.target.value))}
          >
            {itemsPerPageOptions.map((option, index) => (
              <option key={index} value={option}>
                {option}
              </option>
            ))}
          </select>
        </div>
        {Array.from({ length: totalPages }, (_, index) => index + 1).map(
          (page) => (
            <button
              key={page}
              onClick={() => handlePageChange(page)}
              className={currentPage === page ? 'active' : ''}
            >
              {page}
            </button>
          ),
        )}
      </div>
    </div>
  );
};

export default Table;
