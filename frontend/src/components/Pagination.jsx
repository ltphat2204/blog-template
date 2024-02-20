const Pagination = ({ currentPage, totalPages, onPageChange }) => {
    // Calculate start and end nodes to display
    let startNode = Math.max(1, currentPage - 2);
    let endNode = Math.min(totalPages, startNode + 4);

    if (endNode - startNode < 4) {
        startNode = Math.max(1, endNode - 4);
    }

    const pageNumbers = [];
    for (let i = startNode; i <= endNode; i++) {
        pageNumbers.push(i);
    }

    const handlePageChange = (page) => {
        if (page !== currentPage && page >= 1 && page <= totalPages) {
            onPageChange(page);
        }
    };

    return (
        <div className='flex justify-end items-center w-full'>
            <button
                disabled={currentPage === 1}
                onClick={() => handlePageChange(currentPage - 1)}
                className='mx-2 py-1 px-2 rounded bg-gray-200'
            >
                Previous
            </button>
            {pageNumbers.map((pageNumber) => (
                <button
                    key={pageNumber}
                    onClick={() => handlePageChange(pageNumber)}
                    style={{ fontWeight: pageNumber === currentPage ? 'bold' : 'normal' }}
                    className='mx-1 py-1 px-2 rounded bg-gray-200'
                >
                    {pageNumber}
                </button>
            ))}
            <button
                disabled={currentPage === totalPages}
                onClick={() => handlePageChange(currentPage + 1)}
                className='mx-2 py-1 px-2 rounded bg-gray-200'
            >
                Next
            </button>
        </div>
    );
};

export default Pagination;
