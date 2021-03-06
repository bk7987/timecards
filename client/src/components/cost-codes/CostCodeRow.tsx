import React from 'react';
import { CostCodeHours } from '../../types';
import { calcHours, formatDate, renderHours } from '../../utils';
import { Tooltip } from '../Tooltip';

interface CostCodeRowProps extends React.HTMLAttributes<HTMLDivElement> {
  costCode: string;
  description?: string;
  payClass?: string;
  hours: CostCodeHours[];
  dates: Date[];
}

export const CostCodeRow: React.FC<CostCodeRowProps> = ({
  costCode,
  hours,
  dates,
  description = '',
  payClass = '',
  className = '',
  ...props
}) => {
  const getHours = (date: Date) => {
    const [hourSet, tagCodes] = calcHours(
      hours.filter(hourSet => hourSet.date === formatDate(date))
    );
    return renderHours(hourSet, tagCodes);
  };

  const getTotalHours = () => {
    const [total] = calcHours(hours);
    return renderHours(total);
  };

  return (
    <div {...props} className={`flex items-center justify-between hover:bg-teal-200 ${className}`}>
      <div className="w-1/5 leading-tight">
        <div className="text-sm">
          <span className="font-bold">{costCode}</span>
          <span className="ml-4">{payClass}</span>
        </div>
        <div className="truncate">
          <Tooltip content={description} delay={[500, null]}>
            <span className="text-xs">{description}</span>
          </Tooltip>
        </div>
      </div>
      {dates.map(date => {
        return (
          <div key={date.toString()} className="w-1/12 flex items-center justify-center text-sm">
            {getHours(date)}
          </div>
        );
      })}
      <div className="w-1/12 flex items-center justify-center text-sm font-bold">
        {getTotalHours()}
      </div>
    </div>
  );
};
