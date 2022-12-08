package adventofcode.year2022;

import adventofcode.BaseTest;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class Day8 extends BaseTest {

	public static final int DAY = 8;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		int SIZE_X = data.size();
		int SIZE_Y = data.get(0).length();
		int[][] field = new int[SIZE_X][SIZE_Y];
		for (int x = 0; x<SIZE_X; x++){
			for (int y = 0; y<SIZE_Y; y++){
				field[x][y]=Integer.parseInt(""+data.get(x).charAt(y));
			}
		}

		Map<String, Integer > visibleMap = new HashMap<>();
		for (int x = 0; x<SIZE_X; x++){
			int current = -1;
			for (int y = 0; y<SIZE_Y; y++){
				if (field[x][y] > current){
					visibleMap.put(x+"-"+y, field[x][y]);
				}
				current = Math.max(current, field[x][y]);
			}
		}
		for (int x = 0; x<SIZE_X; x++){
			int current = -1;
			for (int y = SIZE_Y-1; y>=0; y--){
				if (field[x][y] > current){
					visibleMap.put(x+"-"+y, field[x][y]);
				}
				current = Math.max(current, field[x][y]);

			}
		}
		for (int y = 0; y<SIZE_Y; y++){
			int current = -1;
			for (int x = SIZE_X-1; x>=0; x--){
				if (field[x][y] > current){
					visibleMap.put(x+"-"+y, field[x][y]);
				}
				current = Math.max(current, field[x][y]);
			}
		}
		for (int y = 0; y<SIZE_Y; y++){
			int current = -1;
			for (int x = 0; x<SIZE_X; x++){
				if (field[x][y] > current){
					visibleMap.put(x+"-"+y, field[x][y]);
				}
				current = Math.max(current, field[x][y]);
			}
		}
		System.out.println(visibleMap.size());
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		int SIZE_X = data.size();
		int SIZE_Y = data.get(0).length();
		int[][] field = new int[SIZE_X][SIZE_Y];
		for (int x = 0; x < SIZE_X; x++) {
			for (int y = 0; y < SIZE_Y; y++) {
				field[x][y] = Integer.parseInt("" + data.get(x).charAt(y));
			}
		}
		int currentScore = 0;
		for (int x = 0; x < SIZE_X; x++) {
			for (int y = 0; y < SIZE_Y; y++) {
				currentScore = Math.max(currentScore, countScore(field, x, y, SIZE_X, SIZE_Y));
			}
		}
		System.out.println(currentScore);
	}

	int countScore(int[][] field, int x, int y, int SIZE_X, int SIZE_Y){
		int seeDown = 0;
		for (int i=1; (i+x)<SIZE_X; i++){
			seeDown++;
			if (field[x][y]<=field[x+i][y]){
				break;
			}
		}

		int seeUp = 0;
		for (int i=1; (x-i)>=0; i++){
			seeUp++;
			if (field[x][y]<=field[x-i][y]){
				break;
			}
		}

		int seeRight = 0;
		for (int j=1; (j+y)<SIZE_Y; j++){
			seeRight++;
			if (field[x][y]<=field[x][y+j]){
				break;
			}
		}

		int seeLeft = 0;
		for (int j=1; j<=y; j++){
			seeLeft++;
			if (field[x][y]<=field[x][y-j]){
				break;
			}
		}
		int score = seeDown*seeUp*seeLeft*seeRight;
//		System.out.println("x="+x+ " y=" + y + "  "+field[x][y]+ "   D " +seeDown + "  U " + seeUp + " L "+ seeLeft +" R "+seeRight + "  "+ score);
		return score;
	}
}
