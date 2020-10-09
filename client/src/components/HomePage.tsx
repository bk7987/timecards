import React from 'react';
import { LinkCard } from './LinkCard';

export const HomePage: React.FC = () => {
  return (
    <div className="px-6 py-3">
      <div className="p-6 rounded border border-gray-300">
        <h2 className="text-3xl font-bold">Welcome to Timecards!</h2>
        <p className="text-gray-600">Click on one of the options below to get started.</p>
      </div>
      <div className="mt-10 flex flex-wrap items-stretch">
        <LinkCard
          to="/foreman-overview"
          description="View the weekly timecard status for each foreman."
          title="Foreman Overview"
          linkText="Foreman overview"
          className="mr-6"
        />
        <LinkCard
          to="/employee-overview"
          description="View all employees with hours on submitted timecards for the selected week."
          title="Employee Overview"
          linkText="Employee overview"
          className="mr-6"
        />
      </div>
    </div>
  );
};
