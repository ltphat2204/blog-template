export default function Milestone({ year, content, title }) {
    return (
        <div className="relative min-h-12 flex items-center pl-16 mb-8 before:contents-['*'] before:absolute before:w-1 before:h-full before:bg-gray-500 before:left-6 before:top-8">
            <div className="text-base absolute left-0 top-0 h-12 w-12 flex items-center justify-center bg-gray-500 text-white rounded-full">{year}</div>
            <div className="text-justify">
                <div className="text-xl font-medium">
                    <h3>{title}</h3>
                </div>
                <div className="text-lg text-slate-600">
                    {content}
                </div>
            </div>
        </div>
    )
}