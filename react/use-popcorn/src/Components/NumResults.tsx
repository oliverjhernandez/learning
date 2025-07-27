type NumResultsProps = {
  results: number;
};

const NumResults = (props: NumResultsProps) => {
  return (
    <p className="num-results">
      Found <strong>{props.results}</strong> results
    </p>
  );
};

export default NumResults;
