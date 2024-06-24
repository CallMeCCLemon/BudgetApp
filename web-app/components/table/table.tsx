import {
    Column,
    Table,
    ColumnDef,
    useReactTable,
    getCoreRowModel,
    getFilteredRowModel,
    getPaginationRowModel,
    flexRender,
    RowData,
} from '@tanstack/react-table'

import React from "react";

// Give our default column cell renderer editing superpowers!
export const defaultColumn: Partial<ColumnDef<object>> = {
    cell: ({ getValue, row: { index }, column: { id }, table }) => {
        const initialValue = getValue()
        // We need to keep and update the state of the cell normally
        // eslint-disable-next-line react-hooks/rules-of-hooks
        const [value, setValue] = React.useState(initialValue)

        // When the input is blurred, we'll call our table meta's updateData function
        const onBlur = () => {
            table.options.meta?.updateData(index, id, value)
        }

        // If the initialValue is changed external, sync it up with our state
        // eslint-disable-next-line react-hooks/rules-of-hooks
        React.useEffect(() => {
            setValue(initialValue)
        }, [initialValue])

        return (
            <input
                value={value as string}
        onChange={e => setValue(e.target.value)}
        onBlur={onBlur}
        />
    )
    },
}

export function useSkipper() {
    const shouldSkipRef = React.useRef(true)
    const shouldSkip = shouldSkipRef.current

    // Wrap a function with this to skip a pagination reset temporarily
    const skip = React.useCallback(() => {
        shouldSkipRef.current = false
    }, [])

    React.useEffect(() => {
        shouldSkipRef.current = true
    })

    return [shouldSkip, skip] as const
}

export interface ColumnHeader {
    name: string,
    type: COLUMN_TYPE,
}

export interface TableProps {
    TableHeader: String
    columnNames: ColumnHeader[]
    rowData: object[]
}

export enum COLUMN_TYPE {
    string,
    number,
    currency,
    date
}

export function Filter({
                    column,
                    table,
                }: {
    column: Column<any, any>
    table: Table<any>
}) {
    const firstValue = table
        .getPreFilteredRowModel()
        .flatRows[0]?.getValue(column.id)

    const columnFilterValue = column.getFilterValue()

    return typeof firstValue === 'number' ? (
            <div className="flex space-x-2">
            <input
                type="number"
        value={(columnFilterValue as [number, number])?.[0] ?? ''}
    onChange={e =>
    column.setFilterValue((old: [number, number]) => [
        e.target.value,
        old?.[1],
    ])
}
    placeholder={`Min`}
    className="w-24 border shadow rounded"
    />
    <input
        type="number"
    value={(columnFilterValue as [number, number])?.[1] ?? ''}
    onChange={e =>
    column.setFilterValue((old: [number, number]) => [
        old?.[0],
        e.target.value,
    ])
}
    placeholder={`Max`}
    className="w-24 border shadow rounded"
        />
        </div>
) : (
        <input
            type="text"
    value={(columnFilterValue ?? '') as string}
    onChange={e => column.setFilterValue(e.target.value)}
    placeholder={`Search...`}
    className="w-36 border shadow rounded"
        />
)
}
