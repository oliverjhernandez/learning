import React from 'react'

const DashboardContainer: React.FC<DashboardContainerProps> = ({
  children,
  color,
  className,
}) => {
  return (
    <div className={`${className}`} style={{ backgroundColor: color }}>
      {children}
    </div>
  )
}

interface DashboardContainerProps {
  children: React.ReactNode
  color: string
  className: string
}

export default DashboardContainer
