# go-sql-export

## Export 1 million records from SQL to CSV or Fixed-Length Format Files with high performance

Exporting large volumes of data efficiently from SQL into a CSV or fixed-length format file requires careful consideration of memory management, I/O performance, and database interactions. This article provides an approach to optimize the process in Golang.

![Export data from SQL to CSV or fixed-length format files](https://cdn-images-1.medium.com/max/800/1*IEMXhQXJ0hWZBPL8q2jMNw.png)
### Test Info
- RAM: 12 GB
- Disk: SSD KINGSTON SA400S37240G ATA Device
- Exec File Size: 10M
- Database: PosgreSQL 16
- Total of rows: 1.018.584 rows
- Total of columns: 76 columns

<table><thead><tr>
<td><b>Type</b></td>
<td><b>File Size</b></td>
<td><b>Rows</b></td>
<td><b>RAM</b></td>
<td><b>Disk</b></td>
<td><b>Power Usage</b></td>
<td><b>Duration</b></td>
<td><b>Description</b></td>
</tr></thead><tbody>

<tr>
<td>Fix Length</td>
<td>1.15 GB</td>
<td>1,018,584</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>Very High</td>
<td>1 min 53 sec</td>
<td>Full scan the table</td>
</tr>

<tr>
<td>CSV</td>
<td>0.975 GB</td>
<td>1,018,584</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>Very High</td>
<td>1 min 25 sec</td>
<td>Full scan the table</td>
</tr>

<tr>
<td>Fix Length</td>
<td>1.02 GB</td>
<td>905,408</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>Very High</td>
<td>1 min 39 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>CSV</td>
<td>0.863 GB</td>
<td>905,408</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>Very High</td>
<td>1 min 10 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>Fix Length</td>
<td>0.89 GB</td>
<td>792,232</td>
<td>14 M</td>
<td>9.9 M/s</td>
<td>Very High</td>
<td>1 min 29 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>CSV</td>
<td>0.764 GB</td>
<td>792,232</td>
<td>14 M</td>
<td>9.9 M/s</td>
<td>Very High</td>
<td>1 min</td>
<td>Filter by index on 1 field</td>
</tr>

</tbody></table>

### Batch jobs
Differ from online processing:
- Long time running, often at night, after working hours.
- Non-interactive, often include logic for handling errors
- Large volumes of data

### Challenges
- <b>Data Volume</b>: Handling 1 million records can strain memory and I/O if not optimized.
- <b>Fixed-Length Format</b>: Properly formatting records to fit fixed-length constraints.
- <b>Database Query Performance</b>: Efficiently retrieving records without overloading the database.

### High-Performance Strategy
#### Streaming Data:
- Instead of loading all records into memory, stream records from SQL to the file. This reduces memory consumption and handles large datasets more efficiently.
- Use GO SDK cursor-based query to fetch data in chunks.
#### Buffered Writing:
- Utilize Go’s bufio.Writer to buffer writes to the file. This reduces the number of I/O operations, improving performance.
#### Efficient Queries:
- Fetch data in batches using LIMIT and OFFSET clauses or use a cursor to paginate through the data.
- Avoid fetching all records at once to prevent memory exhaustion.
#### Error Handling:
- Implement robust error handling to manage failures during the export process, such as network interruptions or file write failures.

### Advantages of This Approach
- <b>Memory Efficiency</b>: Streaming data minimizes memory usage, making the export process scalable for large datasets.
- <b>I/O Optimization</b>: Buffered writing reduces the overhead of frequent file I/O operations.

### Disadvantages
- <b>Do not handle Parallel Processing</b>: If implemented, it can greatly accelerate the export process by utilizing multiple cores.

## Conclusion
High-performance batch processing for exporting data from SQL to CSV or fixed-length files requires careful management of database queries, memory usage, and file I/O. By leveraging streaming, batching, and buffered writing in Golang, you can efficiently handle large-scale exports, ensuring optimal performance while maintaining data integrity.