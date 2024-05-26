import React, {useEffect, useState} from "react";

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

    const DEFAULT_ACTIVE_ROW = -1;
    const [activeRow, setActiveRow] = useState(DEFAULT_ACTIVE_ROW)
    const [activeRowValues, setActiveRowValues] = useState({})

    useEffect(() => {

    }, [activeRow])


    const handleRowClick = (idx: number, rowValues: object) => {
        // TODO: If row changes from one active row to another, save previously active row.

        setActiveRow(idx);
        setActiveRowValues(rowValues);
    }

    const handleSave = () => {
        // TODO: Handle actually saving the data.
        setActiveRow(DEFAULT_ACTIVE_ROW);
    }

    const handleRowChange = (header: string, value: any, type: COLUMN_TYPE) => {
        // TODO: Update the state correctly. Currently, data in inputs is not updating as expected... Forms are whack.
        let updatedActiveRow = activeRowValues;
        if (header in updatedActiveRow) {
            // @ts-ignore
            updatedActiveRow[header] = value
        }
        setActiveRowValues(updatedActiveRow)
    }

    const createTableCols = (columnNames: ColumnHeader[]) => {
        const cols: React.JSX.Element[] = columnNames.map((colHeader, idx) => {
            return <th className="bg-gray-200 p-2 text-left border-b border-gray-300" key={idx}>
                {colHeader.name}
            </th>
        });
        return <thead>
        <tr>
            {cols}
        </tr>
        </thead>
    }

    const formatContent = (content: string, type: COLUMN_TYPE) => {
        let formattedContent = content;
        if (type == COLUMN_TYPE.currency) {
            formattedContent = `\$${parseFloat(formattedContent).toFixed(2)}`
        }

        return <div>{formattedContent}</div>;
    }

    const createTableRow = (rowData: any, colNames: ColumnHeader[], trIdx: number): React.JSX.Element => {
        const rowContents = colNames.map((header, idx) => {
            let content = formatContent(rowData[header.name], header.type);

            if (trIdx == activeRow) {
                content = <input
                // @ts-ignore
                    value={activeRowValues[header.name]}
                    onChange={(event) =>
                        handleRowChange(header.name, event.target.value, header.type)}/>;
            }
            return <td className="flex-1 p-2 border-b border-gray-300" key={idx}
                       onClick={() => handleRowClick(trIdx, rowData)}>
                {content}
            </td>
        });

        if (trIdx === activeRow) {
            rowContents.push(<td>
                <button onClick={() => handleSave()}>Save</button>
            </td>)
        }

        return <tr className="flex justify-between w-full" key={trIdx}>
            {rowContents}
        </tr>
    }

    const tableRows: React.JSX.Element[] = props.rowData.map((row, idx) => {
        return createTableRow(row, props.columnNames, idx);
    });

    return (<div>
        <table className="w-full h-full border-collapse flex flex-col" style={{width: "100%"}}>
            {createTableCols(props.columnNames)}
            <tbody className="flex flex-col h-full">
            {tableRows}
            </tbody>
        </table>
    </div>);
}

