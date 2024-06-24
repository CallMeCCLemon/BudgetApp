import React, {useEffect, useState} from "react";

//
import './table.css'
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
} from '@tanstack/react-table';
import {defaultColumn, TableProps, useSkipper, Filter} from "@/components/table/table";

declare module '@tanstack/react-table' {
    interface TableMeta<TData extends RowData> {
        updateData: (rowIndex: number, columnId: string, value: unknown) => void
    }
}

export default function TransactionTable(props: TableProps) {
    const rerender = React.useReducer(() => ({}), {})[1]

    const columns = React.useMemo<ColumnDef<object>[]>(
        () => [
            {
                header: "Transactions",
                footer: props => props.column.id,
                columns: [
                    {
                        accessorKey: 'Title',
                        footer: props => props.column.id,
                    },
                    {
                        accessorKey: 'ID',
                        footer: props => props.column.id,
                    },
                    {
                        accessorKey: "Total",
                        footer: props => props.column.id,
                    },
                    {
                        accessorKey: "Spent",
                        footer: props => props.column.id
                    },
                    {
                        accessorKey: "Allocated",
                        footer: props => props.column.id
                    }
                ],
            },
        ],
        []
    )

    const [data, setData] = React.useState(() => props.rowData)
    const refreshData = () => setData(() => props.rowData)

    const [autoResetPageIndex, skipAutoResetPageIndex] = useSkipper()

    const table = useReactTable({
        data,
        columns,
        defaultColumn,
        getCoreRowModel: getCoreRowModel(),
        getFilteredRowModel: getFilteredRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        autoResetPageIndex,
        // Provide our updateData function to our table meta
        meta: {
            updateData: (rowIndex, columnId, value) => {
                // Skip page index reset until after next rerender
                skipAutoResetPageIndex()
                setData(old =>
                    old.map((row, index) => {
                        if (index === rowIndex) {
                            return {
                                ...old[rowIndex]!,
                                [columnId]: value,
                            }
                        }
                        return row
                    })
                )
            },
        },
        debugTable: true,
    })

    return (
        <div className="p-2">
        <div className="h-2" />
        <table>
        <thead>
        {table.getHeaderGroups().map(headerGroup => (
            <tr key={headerGroup.id}>
                {headerGroup.headers.map(header => {
                    return (
                        <th key={header.id} colSpan={header.colSpan}>
                            {header.isPlaceholder ? null : (
                                <div>
                                    {flexRender(
                                        header.column.columnDef.header,
                                        header.getContext()
                                    )}
                                    {header.column.getCanFilter() ? (
                                        <div>
                                            <Filter column={header.column} table={table} />
                                        </div>
                                    ) : null}
                                </div>
                            )}
                        </th>
                    )
                })}
            </tr>
        ))}
        </thead>
        <tbody>
        {table.getRowModel().rows.map(row => {
            return (
                <tr key={row.id}>
                    {row.getVisibleCells().map(cell => {
                        return (
                            <td key={cell.id}>
                                {flexRender(
                                    cell.column.columnDef.cell,
                                    cell.getContext()
                                )}
                            </td>
                        )
                    })}
                </tr>
            )
        })}
        </tbody>
    </table>
    <div className="h-2" />
    <div className="flex items-center gap-2">
        <button
            className="border rounded p-1"
            onClick={() => table.setPageIndex(0)}
            disabled={!table.getCanPreviousPage()}
        >
            {'<<'}
        </button>
        <button
            className="border rounded p-1"
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
        >
            {'<'}
        </button>
        <button
            className="border rounded p-1"
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
        >
            {'>'}
        </button>
        <button
            className="border rounded p-1"
            onClick={() => table.setPageIndex(table.getPageCount() - 1)}
            disabled={!table.getCanNextPage()}
        >
            {'>>'}
        </button>
        <span className="flex items-center gap-1">
          <div>Page</div>
          <strong>
            {table.getState().pagination.pageIndex + 1} of{' '}
              {table.getPageCount()}
          </strong>
        </span>
        <span className="flex items-center gap-1">
          | Go to page:
          <input
              type="number"
              defaultValue={table.getState().pagination.pageIndex + 1}
              onChange={e => {
                  const page = e.target.value ? Number(e.target.value) - 1 : 0
                  table.setPageIndex(page)
              }}
              className="border p-1 rounded w-16"
          />
        </span>
        <select
            value={table.getState().pagination.pageSize}
            onChange={e => {
                table.setPageSize(Number(e.target.value))
            }}
        >
            {[10, 20, 30, 40, 50].map(pageSize => (
                <option key={pageSize} value={pageSize}>
                    Show {pageSize}
                </option>
            ))}
        </select>
    </div>
    <div>{table.getRowModel().rows.length} Rows</div>
    <div>
        <button onClick={() => rerender()}>Force Rerender</button>
    </div>
    <div>
        <button onClick={() => refreshData()}>Refresh Data</button>
    </div>
</div>
)
}
