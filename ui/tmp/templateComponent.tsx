import React from "react";
import styled from "styled-components";

// Type層
type Props = {
  className?: string;
};

// DOM層
const Component: React.FC<Props> = (props) => {
  return <div className={props.className}></div>;
};

//Style層
const StyledComponent = styled(Component)``;

export default StyledComponent;
