interface ErrorDisplayProps {
    error: string | null;
    onErrorDismiss: () => void;
}

export const ErrorDisplay: React.FC<ErrorDisplayProps> = ({ error, onErrorDismiss }) => {
    if (!error) return null;

    return (
        <div className="bg-red-100 border-l-4 border-red-500 text-red-700 p-4 rounded-lg shadow-md flex justify-between items-center">
            <p className="font-medium flex items-center">
                <p className="w-5 h-5 mr-2" />
                {error}
            </p>
            <button onClick={onErrorDismiss} className="text-red-500 hover:text-red-800 font-bold">
                &times;
            </button>
        </div>
    );
};