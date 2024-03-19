import React from 'react';
import { render } from '@testing-library/react';
import '@testing-library/jest-dom';
import Card from './Card';
describe('Card Component', () => {
    test('renders card with title and value', () => {
        const title = 'Test Title';
        const value = 'Test Value';

        const { getByText } = render(<Card title={ title } value = { value } />);

        const titleElement = getByText(title);
        const valueElement = getByText(value);

        expect(titleElement).toBeInTheDocument();
        expect(valueElement).toBeInTheDocument();
    });
});
