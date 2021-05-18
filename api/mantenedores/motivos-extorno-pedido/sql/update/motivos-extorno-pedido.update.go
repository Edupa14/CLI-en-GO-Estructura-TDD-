CREATE   PROCEDURE [dbo].[NS_TMMOTIVOS_EXTORNO_PEDIDO_U](
	@ID INT,
	@JSON NVARCHAR(MAX)
)
AS
	DECLARE @ErrorMessage NVARCHAR(max),
			@ErrorSeverity INT,
			@ErrorState INT;

BEGIN TRANSACTION
BEGIN TRY
	IF NOT EXISTS(SELECT id FROM TMMOTIVOS_EXTORNO_PEDIDO WITH(NOLOCK) WHERE id = @id)
		THROW 51008, 'No existe este registro', 15;

		WITH HEADER (nombre,codigo) AS (
			SELECT
				nombre,codigo
			FROM OPENJSON(@JSON)
			WITH (
				nombre VARCHAR(250),codigo VARCHAR(15)
			)
		)
		UPDATE [target] SET
                [target].nombre= [source].nombre,[target].codigo= [source].codigo
        FROM HEADER [source]
        INNER JOIN TMMOTIVOS_EXTORNO_PEDIDO [target] WITH(HOLDLOCK) ON [target].id = @ID;
		
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