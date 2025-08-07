type ErrorMessageType = {
  message: string;
};

const ErrorMessage = ({ message }: ErrorMessageType) => {
  return (
    <div>
      <p className="error">{message}</p>
    </div>
  );
};

export default ErrorMessage;
