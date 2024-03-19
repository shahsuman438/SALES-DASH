import React from 'react';

interface CardProps {
  title: string;
  value: string;
}

const Card: React.FC<CardProps> = ({ title, value }) => {
  return (
    <div className='card'>
      <h3 className='card__title'>{title}</h3>
      <p className='card__value'>{value}</p>
    </div>
  );
};

export default Card;
