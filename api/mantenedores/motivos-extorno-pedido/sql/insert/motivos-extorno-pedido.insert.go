CREATE   PROCEDURE [dbo].[NS_TMMOTIVOS_EXTORNO_PEDIDO_I](
	@JSON VARCHAR(MAX)
)
AS
DECLARE @ErrorMessage NVARCHAR(max),
		@ErrorSeverity INT,
		@ErrorState INT,
		@id smallint;

DECLARE @temptable table(
	id smallint
)
BEGIN TRANSACTION
	BEGIN TRY;
        WITH MOTIVOSEXTORNOPEDIDO (nombre,codigo) AS (
			SELECT
				nombre,codigo
			FROM OPENJSON(@JSON)
			WITH (
				nombre VARCHAR(250),codigo VARCHAR(15)
			)
		)
		INSERT INTO TMMOTIVOS_EXTORNO_PEDIDO(
				nombre,codigo
		)
		OUTPUT INSERTED.id INTO @temptable(id)
		SELECT nombre,codigo
        		FROM MOTIVOSEXTORNOPEDIDO
		SELECT @id = id FROM @temptable ;

		SELECT @id as id

		COMMIT TRANSACTION;

	END TRY
BEGIN CATCH
	ROLLBACK TRANSACTION;
	SELECT
		@ErrorMessage='Ocurrio un Error: '+ERROR_MESSAGE() + ' LN. ' + CONVERT(NVARCHAR(255), ERROR_LINE() ) + '.',
		@ErrorSeverity=ERROR_SEVERITY(),
		@ErrorState=ERROR_STATE();
	THROW 51003, @ErrorMessage, 10
END CATCH
