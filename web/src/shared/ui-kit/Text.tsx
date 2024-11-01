import cn from "classnames";
import { ComponentProps } from "react";

interface Props extends React.HTMLAttributes<HTMLSpanElement> {
  type?: 'default' | 'sm-1' | 'sm-2' | 'header' | 'subheader' | 'section' | 'large',
  children?: React.ReactNode
}

export type TextProps = ComponentProps<typeof Text>

export const Text = ({ type = 'default', children, className, ...attrs }: Props) => {
  return (
    <span
      {...attrs}
      className={cn(
        className,
        {
          'text-[0.9rem] leading-[1.125rem] -tracking-[1%]': type === 'default',
          'text-[1.5rem] leading-[1.8rem] -tracking-[1%] font-medium': type === 'header',
          'text-[1.25rem] leading-[1.5rem] -tracking-[1%] font-medium': type === 'subheader',
          'text-[0.8rem] leading-[1rem]': type === 'sm-1',
          'text-[0.7rem] leading-[0.8rem]': type === 'sm-2',
          'text-[0.7rem] leading-[0.8rem] tracking-[3%] uppercase font-medium': type === 'section',
        }
      )}
    >
      {children}
    </span>
  )
}