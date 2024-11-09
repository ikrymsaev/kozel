import { ButtonHTMLAttributes, MouseEvent, PropsWithChildren } from "react"
import cn from "classnames"
import { SpinnerIcon } from "../icons/SpinnerIcon"

interface Props extends PropsWithChildren<ButtonHTMLAttributes<HTMLButtonElement>> {
  size?: "xs" | "s" | "m"
  color?: keyof typeof bgStyles,
  loading?: boolean // TODO: implement
  mode?: "bezeled"
  before?: React.ReactNode
  fill?: boolean
  badge?: boolean
  disabledType?: "inactive"
}

export const Button = ({
  children, loading, className, before = null, size = 's', color = 'accent', mode, fill, badge, onClick, disabledType,
  ...props
}: Props) => {

  const handleClick = (e: MouseEvent<HTMLButtonElement, globalThis.MouseEvent>) => {
    if (loading) return;
    onClick?.(e)
  }

  return (
    <button
      onClick={handleClick}
      className={cn(
        className,
        props.disabled && disabledType === 'inactive' && 'disabled:bg-[#F1EFEF] text-lightGray',
        ...baseStyles,
        bgStyles[color],
        sizeStyles[size],
        mode === 'bezeled' ? 'text-link' : color === 'white' ? 'text-accent' : 'text-btnText',
        mode === 'bezeled' && 'bg-bezeled',
        color !== 'transparent' && 'disabled:bg-inactive ',
        fill ? 'w-full' : 'w-fit',
      )}
      {...props}
    >
      {!loading && before}
      {loading && <SpinnerIcon width={24} height={24} className="animate-spin" />}
      {!loading && children}
      {badge && <div className="absolute right-0 top-0 h-2 w-2 rounded-full bg-stopRed" />}
    </button>
  )
}

const baseStyles = [
  "flex items-center justify-center gap-[6px] -tracking-[0.3px] relative",
  "cursor-pointer hover:opacity-90",
  "active:scale-[99%]",
  "disabled:cursor-default disabled:hover:opacity-100",
]
const sizeStyles = {
  'xs': "h-[1.5rem] px-4 rounded-[20px] min-w-[5rem] text-[1rem] leading-[1rem] font-bold",
  's': "h-[2rem] px-3 rounded-[20px] min-w-[5rem] text-sm font-bold",
  'm': "h-[3rem] px-4 rounded-[10px] text-[1.2rem] leading-[1.2rem] font-semibold",
}
const bgStyles = {
  accent: "bg-accent",
  success: "bg-success",
  transparent: "bg-transparent",
  purple: "bg-purple",
  white: "bg-white",
  gold: "bg-gold",
}