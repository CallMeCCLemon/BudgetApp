import React, {useState} from "react";
import {list} from "postcss";

export interface ColumnHeader {
    name: string,
    type: COLUMN_TYPE,
}

export interface TableProps {
    columnNames: ColumnHeader[]
    rowData: object[]
}

export enum COLUMN_TYPE {
    string,
    number,
    currency,
    date
}

export default function Table(props: TableProps) {
    const [activeRow, setActiveRow] = useState(2)
    const [activeRowValues, setActiveRowValues] = useState({})

    const handleRowClick = (event: Event) => {
        // setActiveRow(event.target)
        // TODO: Look up how to set the active row upon a click event
    }

    const createTableCols = (columnNames: ColumnHeader[]) => {
        const cols: React.JSX.Element[] = columnNames.map((colHeader, idx) => {
            return <th className="bg-gray-200 p-2 text-left border-b border-gray-300" key={idx}>
                { colHeader.name }
            </th>
        });
        return <thead>
            <tr>
            {cols}
            </tr>
        </thead>
    }

    const createTableRow = (rowData: any, colNames: ColumnHeader[], trIdx: number): React.JSX.Element => {
        const rowContents = colNames.map((header, idx) => {
            let content = rowData[header.name];
            if (header.type == COLUMN_TYPE.currency) {
                content = `\$${content.toFixed(2)}`
            }
            if (trIdx == activeRow) {
                content = <input value={content}/>;
            }
            return <td className="flex-1 p-2 border-b border-gray-300" key={idx}>
                {/*onClick={handleRowClick}>*/}
                { content }
            </td>
        });

        if (trIdx === activeRow) {
            rowContents.push(<td><button>Save</button></td>)
        }

        console.log(trIdx);

        return <tr className="flex justify-between w-full" key={trIdx}>
            { rowContents }
        </tr>
    }

    const tableRows: React.JSX.Element[] = props.rowData.map((row, idx) => {
        return createTableRow(row, props.columnNames, idx);
    });

    return (<div>
        <table className="w-full h-full border-collapse flex flex-col" style={{width: "100%"}}>
            { createTableCols(props.columnNames) }
            <tbody className="flex flex-col h-full">
            { tableRows }
            </tbody>
        </table>
    </div>);
}

