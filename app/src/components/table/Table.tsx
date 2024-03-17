import React, { useState } from 'react';

interface Props {
    data: any[];
    itemsPerPage: number;
}

const Table: React.FC<Props> = ({ data, itemsPerPage }) => {
    const [currentPage, setCurrentPage] = useState(1);

    const startIndex = (currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    const currentItems = data.slice(startIndex, endIndex);
    const totalPages = Math.ceil(data.length / itemsPerPage);

    const handlePageChange = (pageNumber: number) => {
        setCurrentPage(pageNumber);
    };

    // Extract headers dynamically from the first item in the data array
    const headers = Object.keys(data[0] || {});

    return (
        <div className="table-container p-1 pt-4">
            <table className="table">
                <thead>
                    <tr>
                        {headers.map((header, index) => (
                            <th key={index}>{header.toLocaleUpperCase()}</th>
                        ))}
                    </tr>
                </thead>
                <tbody>
                    {currentItems.map((item: any, index: number) => (
                        <tr key={index}>
                            {headers.map((header, index) => (
                                <td key={index}>{item[header]}</td>
                            ))}
                        </tr>
                    ))}
                </tbody>
            </table>

            <div className="pagination">
                {/* <select name="" id="">
                    <option hidden selected>Items per page</option>
                    <option value="5">5</option>
                    <option value="10">10</option>
                    <option value="20">20</option>
                    <option value="100">100</option>
                    <option value="200">200</option>
                </select> */}
                {Array.from({ length: totalPages }, (_, index) => index + 1).map((page) => (
                    <button
                        key={page}
                        onClick={() => handlePageChange(page)}
                        className={currentPage === page ? 'active' : ''}
                    >
                        {page}
                    </button>
                ))}
            </div>
        </div>
    );
};

export default Table;
