import { useState } from 'react';

const Dropdown = ({ options, onSelect }) => {
    const [selectedOption, setSelectedOption] = useState('');

    const handleSelect = (e) => {
        const selectedId = e.target.value;
        setSelectedOption(selectedId);
        onSelect(selectedId);
    };

    return (
        <select value={selectedOption} onChange={handleSelect} className='bg-gray-200 px-2 py-1 rounded'>
            <option value="">
                All
            </option>
            {options && options.map((option) => (
                <option key={option.id} value={option.id}>
                    {option.name}
                </option>
            ))}
        </select>
    );
};

export default Dropdown;
